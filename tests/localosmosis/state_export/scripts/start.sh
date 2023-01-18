#!/bin/sh
set -e 
set -o pipefail

MOKITA_HOME=$HOME/.mokitad
CONFIG_FOLDER=$MOKITA_HOME/config

DEFAULT_MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"
DEFAULT_CHAIN_ID="localmokita"
DEFAULT_MONIKER="val"

# Override default values with environment variables
MNEMONIC=${MNEMONIC:-$DEFAULT_MNEMONIC}
CHAIN_ID=${CHAIN_ID:-$DEFAULT_CHAIN_ID}
MONIKER=${MONIKER:-$DEFAULT_MONIKER}

install_prerequisites () {
    apk add -q --no-cache \
        dasel \
        python3 \
        py3-pip
}

edit_config () {

    # Remove seeds
    dasel put string -f $CONFIG_FOLDER/config.toml '.p2p.seeds' ''

    # Disable fast_sync
    dasel put bool -f $CONFIG_FOLDER/config.toml '.fast_sync' 'false'

    # Expose the rpc
    dasel put string -f $CONFIG_FOLDER/config.toml '.rpc.laddr' "tcp://0.0.0.0:26657"
}

if [[ ! -d $CONFIG_FOLDER ]]
then

    install_prerequisites

    echo "Chain ID: $CHAIN_ID"
    echo "Moniker:  $MONIKER"

    echo $MNEMONIC | mokitad init -o --chain-id=$CHAIN_ID --home $MOKITA_HOME --recover $MONIKER 2> /dev/null
    echo $MNEMONIC | mokitad keys add my-key --recover --keyring-backend test > /dev/null 2>&1

    ACCOUNT_PUBKEY=$(mokitad keys show --keyring-backend test my-key --pubkey | dasel -r json '.key' --plain)
    ACCOUNT_ADDRESS=$(mokitad keys show -a --keyring-backend test my-key --bech acc)

    VALIDATOR_PUBKEY_JSON=$(mokitad tendermint show-validator --home $MOKITA_HOME)
    VALIDATOR_PUBKEY=$(echo $VALIDATOR_PUBKEY_JSON | dasel -r json '.key' --plain)
    VALIDATOR_HEX_ADDRESS=$(mokitad debug pubkey $VALIDATOR_PUBKEY_JSON 2>&1 --home $MOKITA_HOME | grep Address | cut -d " " -f 2)
    VALIDATOR_ACCOUNT_ADDRESS=$(mokitad debug addr $VALIDATOR_HEX_ADDRESS 2>&1  --home $MOKITA_HOME | grep Acc | cut -d " " -f 3)
    VALIDATOR_OPERATOR_ADDRESS=$(mokitad debug addr $VALIDATOR_HEX_ADDRESS 2>&1  --home $MOKITA_HOME | grep Val | cut -d " " -f 3)
    VALIDATOR_CONSENSUS_ADDRESS=$(mokitad debug bech32-convert $VALIDATOR_OPERATOR_ADDRESS -p mokivalcons  --home $MOKITA_HOME 2>&1)

    python3 -u testnetify.py \
    -i /mokita/state_export.json \
    -o $CONFIG_FOLDER/genesis.json \
    -c $CHAIN_ID \
    --validator-hex-address $VALIDATOR_HEX_ADDRESS \
    --validator-operator-address $VALIDATOR_OPERATOR_ADDRESS \
    --validator-consensus-address $VALIDATOR_CONSENSUS_ADDRESS \
    --validator-pubkey $VALIDATOR_PUBKEY \
    --account-pubkey $ACCOUNT_PUBKEY \
    --account-address $ACCOUNT_ADDRESS \
    --prune-ibc

    edit_config
fi

mokitad start --home $MOKITA_HOME --x-crisis-skip-assert-invariants
