#!/bin/bash
set -o errexit -o nounset -o pipefail

PASSWORD=${PASSWORD:-1234567890}
STAKE=${STAKE_TOKEN:-ufury}
FEE=${FEE_TOKEN:-ufury}
CHAIN_ID=${CHAIN_ID:-testing}
MONIKER=${MONIKER:-node001}

gridiron init --chain-id "$CHAIN_ID" "$MONIKER"
# staking/governance token is hardcoded in config, change this
## OSX requires: -i.
sed -i. "s/\"stake\"/\"$STAKE\"/" "$HOME"/.gridiron/config/genesis.json
for x in validator bootstrap-account; do
  if ! gridiron keys show $x; then
    (
      echo "$PASSWORD"
      echo "$PASSWORD"
    ) | gridiron keys add "$x"
  fi
done

# set date based on OS
end_time='unknown'
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  end_time=$(date -d "+10 years" +%s)
elif [[ "$OSTYPE" == "darwin"* ]]; then
  end_time=$(date -v+10y +%s)
fi

# hardcode the account for this instance
echo "$PASSWORD" | gridiron add-genesis-account bootstrap-account "1000000000$STAKE"
echo "$PASSWORD" | gridiron add-genesis-account validator "1001000000$STAKE" --vesting-amount="1000000000$STAKE" --vesting-end-time="$end_time"


# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo "$addr"
  gridiron add-genesis-account "$addr" "1000000000$STAKE"
done

## POE setup
# set bootstrap-account address (temporary)

# set engagement points
content=$(cat "$HOME"/.gridiron/config/genesis.json | jq  ".app_state.poe.seed_contracts.engagement |= . + [{\"address\":\"$(echo "$PASSWORD" | gridiron keys show -a validator)\",\"points\":\"100\"}]")
# set oversight community
content=$(echo "$content" | jq  ".app_state.poe.seed_contracts.oversight_community_members |= . + [\"$(echo "$PASSWORD" | gridiron keys show -a bootstrap-account)\"]")
# set arbiter
content=$(echo "$content" | jq  ".app_state.poe.seed_contracts.arbiter_pool_members |= . + [\"$(echo "$PASSWORD" | gridiron keys show -a bootstrap-account)\"]")
# set system admin
content=$(echo "$content" | jq  ".app_state.poe.seed_contracts.bootstrap_account_address |= \"$(echo "$PASSWORD" | gridiron keys show -a bootstrap-account)\"")
# set min fee
content=$(echo "$content" | jq  ".app_state.globalfee.params.minimum_gas_prices |= [{\"denom\":\"$STAKE\",\"amount\":\"0.001\"}]")

mv "$HOME"/.gridiron/config/genesis.json  "$HOME"/.gridiron/config/genesis.json_old
echo "$content" > "$HOME"/.gridiron/config/genesis.json

# submit a genesis validator tx
## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
(
  echo "$PASSWORD"
  echo "$PASSWORD"
  echo "$PASSWORD"
) | gridiron gentx validator "0$STAKE" "250000000$STAKE" --chain-id="$CHAIN_ID" --amount="0$STAKE" --vesting-amount="250000000$STAKE" --fees="2000$STAKE"

gridiron collect-gentxs
