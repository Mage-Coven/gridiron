<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [mage-coven/globalfee/v1beta1/genesis.proto](#mage-coven/globalfee/v1beta1/genesis.proto)
    - [GenesisState](#mage-coven.globalfee.v1beta1.GenesisState)
    - [Params](#mage-coven.globalfee.v1beta1.Params)
  
- [mage-coven/globalfee/v1beta1/query.proto](#mage-coven/globalfee/v1beta1/query.proto)
    - [QueryMinimumGasPricesRequest](#mage-coven.globalfee.v1beta1.QueryMinimumGasPricesRequest)
    - [QueryMinimumGasPricesResponse](#mage-coven.globalfee.v1beta1.QueryMinimumGasPricesResponse)
  
    - [Query](#mage-coven.globalfee.v1beta1.Query)
  
- [mage-coven/poe/v1beta1/poe.proto](#mage-coven/poe/v1beta1/poe.proto)
    - [Params](#mage-coven.poe.v1beta1.Params)
  
    - [PoEContractType](#mage-coven.poe.v1beta1.PoEContractType)
  
- [mage-coven/poe/v1beta1/genesis.proto](#mage-coven/poe/v1beta1/genesis.proto)
    - [ArbiterPoolContractConfig](#mage-coven.poe.v1beta1.ArbiterPoolContractConfig)
    - [CommunityPoolContractConfig](#mage-coven.poe.v1beta1.CommunityPoolContractConfig)
    - [EngagementContractConfig](#mage-coven.poe.v1beta1.EngagementContractConfig)
    - [GenesisState](#mage-coven.poe.v1beta1.GenesisState)
    - [ImportDump](#mage-coven.poe.v1beta1.ImportDump)
    - [MixerContractConfig](#mage-coven.poe.v1beta1.MixerContractConfig)
    - [MixerContractConfig.Sigmoid](#mage-coven.poe.v1beta1.MixerContractConfig.Sigmoid)
    - [OversightCommitteeContractConfig](#mage-coven.poe.v1beta1.OversightCommitteeContractConfig)
    - [PoEContract](#mage-coven.poe.v1beta1.PoEContract)
    - [SeedContracts](#mage-coven.poe.v1beta1.SeedContracts)
    - [StakeContractConfig](#mage-coven.poe.v1beta1.StakeContractConfig)
    - [TG4Member](#mage-coven.poe.v1beta1.TG4Member)
    - [ValidatorVotingContractConfig](#mage-coven.poe.v1beta1.ValidatorVotingContractConfig)
    - [ValsetContractConfig](#mage-coven.poe.v1beta1.ValsetContractConfig)
    - [VotingRules](#mage-coven.poe.v1beta1.VotingRules)
  
- [mage-coven/poe/v1beta1/query.proto](#mage-coven/poe/v1beta1/query.proto)
    - [QueryContractAddressRequest](#mage-coven.poe.v1beta1.QueryContractAddressRequest)
    - [QueryContractAddressResponse](#mage-coven.poe.v1beta1.QueryContractAddressResponse)
    - [QueryUnbondingPeriodRequest](#mage-coven.poe.v1beta1.QueryUnbondingPeriodRequest)
    - [QueryUnbondingPeriodResponse](#mage-coven.poe.v1beta1.QueryUnbondingPeriodResponse)
    - [QueryValidatorDelegationRequest](#mage-coven.poe.v1beta1.QueryValidatorDelegationRequest)
    - [QueryValidatorDelegationResponse](#mage-coven.poe.v1beta1.QueryValidatorDelegationResponse)
    - [QueryValidatorEngagementRewardRequest](#mage-coven.poe.v1beta1.QueryValidatorEngagementRewardRequest)
    - [QueryValidatorEngagementRewardResponse](#mage-coven.poe.v1beta1.QueryValidatorEngagementRewardResponse)
    - [QueryValidatorOutstandingRewardRequest](#mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardRequest)
    - [QueryValidatorOutstandingRewardResponse](#mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardResponse)
    - [QueryValidatorUnbondingDelegationsRequest](#mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsRequest)
    - [QueryValidatorUnbondingDelegationsResponse](#mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsResponse)
  
    - [Query](#mage-coven.poe.v1beta1.Query)
  
- [mage-coven/poe/v1beta1/tx.proto](#mage-coven/poe/v1beta1/tx.proto)
    - [MsgCreateValidator](#mage-coven.poe.v1beta1.MsgCreateValidator)
    - [MsgCreateValidatorResponse](#mage-coven.poe.v1beta1.MsgCreateValidatorResponse)
    - [MsgDelegate](#mage-coven.poe.v1beta1.MsgDelegate)
    - [MsgDelegateResponse](#mage-coven.poe.v1beta1.MsgDelegateResponse)
    - [MsgUndelegate](#mage-coven.poe.v1beta1.MsgUndelegate)
    - [MsgUndelegateResponse](#mage-coven.poe.v1beta1.MsgUndelegateResponse)
    - [MsgUpdateValidator](#mage-coven.poe.v1beta1.MsgUpdateValidator)
    - [MsgUpdateValidatorResponse](#mage-coven.poe.v1beta1.MsgUpdateValidatorResponse)
  
    - [Msg](#mage-coven.poe.v1beta1.Msg)
  
- [mage-coven/twasm/v1beta1/contract_extension.proto](#mage-coven/twasm/v1beta1/contract_extension.proto)
    - [RegisteredPrivilege](#mage-coven.twasm.v1beta1.RegisteredPrivilege)
    - [GridironContractDetails](#mage-coven.twasm.v1beta1.GridironContractDetails)
  
- [mage-coven/twasm/v1beta1/genesis.proto](#mage-coven/twasm/v1beta1/genesis.proto)
    - [Contract](#mage-coven.twasm.v1beta1.Contract)
    - [CustomModel](#mage-coven.twasm.v1beta1.CustomModel)
    - [GenesisState](#mage-coven.twasm.v1beta1.GenesisState)
    - [KVModel](#mage-coven.twasm.v1beta1.KVModel)
  
- [mage-coven/twasm/v1beta1/proposal.proto](#mage-coven/twasm/v1beta1/proposal.proto)
    - [DemotePrivilegedContractProposal](#mage-coven.twasm.v1beta1.DemotePrivilegedContractProposal)
    - [PromoteToPrivilegedContractProposal](#mage-coven.twasm.v1beta1.PromoteToPrivilegedContractProposal)
  
- [mage-coven/twasm/v1beta1/query.proto](#mage-coven/twasm/v1beta1/query.proto)
    - [QueryContractsByPrivilegeTypeRequest](#mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeRequest)
    - [QueryContractsByPrivilegeTypeResponse](#mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeResponse)
    - [QueryPrivilegedContractsRequest](#mage-coven.twasm.v1beta1.QueryPrivilegedContractsRequest)
    - [QueryPrivilegedContractsResponse](#mage-coven.twasm.v1beta1.QueryPrivilegedContractsResponse)
  
    - [Query](#mage-coven.twasm.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="mage-coven/globalfee/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/globalfee/v1beta1/genesis.proto



<a name="mage-coven.globalfee.v1beta1.GenesisState"></a>

### GenesisState
GenesisState - initial state of module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#mage-coven.globalfee.v1beta1.Params) |  | Params of this module |






<a name="mage-coven.globalfee.v1beta1.Params"></a>

### Params
Params defines the set of module parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minimum_gas_prices` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | Minimum stores the minimum gas price(s) for all TX on the chain. When multiple coins are defined then they are accepted alternatively. The list must be sorted by denoms asc. No duplicate denoms or zero amount values allowed. For more information see https://docs.cosmos.network/master/modules/auth/01_concepts.html |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/globalfee/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/globalfee/v1beta1/query.proto



<a name="mage-coven.globalfee.v1beta1.QueryMinimumGasPricesRequest"></a>

### QueryMinimumGasPricesRequest
QueryMinimumGasPricesRequest is the request type for the
Query/MinimumGasPrices RPC method.






<a name="mage-coven.globalfee.v1beta1.QueryMinimumGasPricesResponse"></a>

### QueryMinimumGasPricesResponse
QueryMinimumGasPricesResponse is the response type for the
Query/MinimumGasPrices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minimum_gas_prices` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="mage-coven.globalfee.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `MinimumGasPrices` | [QueryMinimumGasPricesRequest](#mage-coven.globalfee.v1beta1.QueryMinimumGasPricesRequest) | [QueryMinimumGasPricesResponse](#mage-coven.globalfee.v1beta1.QueryMinimumGasPricesResponse) |  | GET|/gridiron/globalfee/v1beta1/minimum_gas_prices|

 <!-- end services -->



<a name="mage-coven/poe/v1beta1/poe.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/poe/v1beta1/poe.proto



<a name="mage-coven.poe.v1beta1.Params"></a>

### Params
Params defines the parameters for the PoE module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `historical_entries` | [uint32](#uint32) |  | HistoricalEntries is the number of historical entries to persist. |
| `initial_val_engagement_points` | [uint64](#uint64) |  | InitialValEngagementPoints defines the number of engagement for any new validator joining post genesis |
| `min_delegation_amounts` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | MinDelegationAmount defines the minimum amount a post genesis validator needs to self delegate to receive any engagement points. One must be exceeded. No minimum condition set when empty. |





 <!-- end messages -->


<a name="mage-coven.poe.v1beta1.PoEContractType"></a>

### PoEContractType
PoEContractType type of PoE contract

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNDEFINED | 0 |  |
| STAKING | 1 |  |
| VALSET | 2 |  |
| ENGAGEMENT | 3 |  |
| MIXER | 4 |  |
| DISTRIBUTION | 5 |  |
| OVERSIGHT_COMMUNITY | 6 |  |
| OVERSIGHT_COMMUNITY_PROPOSALS | 7 |  |
| COMMUNITY_POOL | 8 |  |
| VALIDATOR_VOTING | 9 |  |
| ARBITER_POOL | 10 |  |
| ARBITER_POOL_VOTING | 11 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/poe/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/poe/v1beta1/genesis.proto



<a name="mage-coven.poe.v1beta1.ArbiterPoolContractConfig"></a>

### ArbiterPoolContractConfig
ArbiterPoolContractConfig initial setup config for the trusted circle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | Name of TRUSTED_CIRCLE |
| `escrow_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | EscrowAmount The required escrow amount, in the default denom (ufury) |
| `voting_rules` | [VotingRules](#mage-coven.poe.v1beta1.VotingRules) |  | VotingRules rules for the tally |
| `deny_list_contract_address` | [string](#string) |  | DenyListContractAddress is an optional cw4 contract with list of addresses denied to be part of TrustedCircle |
| `dispute_cost` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | DisputeCost The required dispute amount, in the default denom (ufury) |
| `waiting_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="mage-coven.poe.v1beta1.CommunityPoolContractConfig"></a>

### CommunityPoolContractConfig
CommunityPoolContractConfig initial setup config for the contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_rules` | [VotingRules](#mage-coven.poe.v1beta1.VotingRules) |  | VotingRules rules for the tally |






<a name="mage-coven.poe.v1beta1.EngagementContractConfig"></a>

### EngagementContractConfig
EngagementContractConfig initial setup config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `halflife` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="mage-coven.poe.v1beta1.GenesisState"></a>

### GenesisState
GenesisState - initial state of module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#mage-coven.poe.v1beta1.Params) |  | params defines all the parameter of the module |
| `seed_contracts` | [SeedContracts](#mage-coven.poe.v1beta1.SeedContracts) |  | SeedContracts stores and instantiates the Proof of Engagement contracts on the chain. |
| `import_dump` | [ImportDump](#mage-coven.poe.v1beta1.ImportDump) |  | ImportDump restores the state from an exported state genesis |






<a name="mage-coven.poe.v1beta1.ImportDump"></a>

### ImportDump
ImportDump has all module data for non seed mode.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [PoEContract](#mage-coven.poe.v1beta1.PoEContract) | repeated | Contracts PoE contract addresses and types |






<a name="mage-coven.poe.v1beta1.MixerContractConfig"></a>

### MixerContractConfig
Mixer contract setup. See
https://github.com/mage-coven/poe-contracts/tree/main/contracts/tg4-mixer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sigmoid` | [MixerContractConfig.Sigmoid](#mage-coven.poe.v1beta1.MixerContractConfig.Sigmoid) |  | Sigmoid returns a sigmoid-like value of staked amount times engagement points. See the Proof-of-Engagement white-paper for details. |






<a name="mage-coven.poe.v1beta1.MixerContractConfig.Sigmoid"></a>

### MixerContractConfig.Sigmoid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_points` | [uint64](#uint64) |  |  |
| `p` | [string](#string) |  |  |
| `s` | [string](#string) |  |  |






<a name="mage-coven.poe.v1beta1.OversightCommitteeContractConfig"></a>

### OversightCommitteeContractConfig
OversightCommitteeContractConfig initial setup config for the trusted circle


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | Name of TRUSTED_CIRCLE |
| `escrow_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | EscrowAmount The required escrow amount, in the default denom (ufury) |
| `voting_rules` | [VotingRules](#mage-coven.poe.v1beta1.VotingRules) |  | VotingRules rules for the tally |
| `deny_list_contract_address` | [string](#string) |  | DenyListContractAddress is an optional cw4 contract with list of addresses denied to be part of TrustedCircle |






<a name="mage-coven.poe.v1beta1.PoEContract"></a>

### PoEContract
PoEContract address and type information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_type` | [PoEContractType](#mage-coven.poe.v1beta1.PoEContractType) |  | ContractType type. |
| `address` | [string](#string) |  | Address is the bech32 address string |






<a name="mage-coven.poe.v1beta1.SeedContracts"></a>

### SeedContracts
SeedContracts contains the contract configuration and group members to setup
all PoE contracts on chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gen_txs` | [bytes](#bytes) | repeated | GenTxs defines the genesis transactions to create a validator. |
| `bootstrap_account_address` | [string](#string) |  | BootstrapAccountAddress single address that is set as bootstrap-account for the PoE contracts in seed mode. |
| `engagement` | [TG4Member](#mage-coven.poe.v1beta1.TG4Member) | repeated | Engagement weighted members of the engagement group. Validators should be in here. |
| `stake_contract_config` | [StakeContractConfig](#mage-coven.poe.v1beta1.StakeContractConfig) |  |  |
| `valset_contract_config` | [ValsetContractConfig](#mage-coven.poe.v1beta1.ValsetContractConfig) |  |  |
| `engagement_contract_config` | [EngagementContractConfig](#mage-coven.poe.v1beta1.EngagementContractConfig) |  |  |
| `bond_denom` | [string](#string) |  | BondDenom defines the bondable coin denomination. |
| `oversight_committee_contract_config` | [OversightCommitteeContractConfig](#mage-coven.poe.v1beta1.OversightCommitteeContractConfig) |  |  |
| `community_pool_contract_config` | [CommunityPoolContractConfig](#mage-coven.poe.v1beta1.CommunityPoolContractConfig) |  |  |
| `validator_voting_contract_config` | [ValidatorVotingContractConfig](#mage-coven.poe.v1beta1.ValidatorVotingContractConfig) |  |  |
| `oversight_community_members` | [string](#string) | repeated |  |
| `arbiter_pool_members` | [string](#string) | repeated |  |
| `arbiter_pool_contract_config` | [ArbiterPoolContractConfig](#mage-coven.poe.v1beta1.ArbiterPoolContractConfig) |  |  |
| `mixer_contract_config` | [MixerContractConfig](#mage-coven.poe.v1beta1.MixerContractConfig) |  |  |






<a name="mage-coven.poe.v1beta1.StakeContractConfig"></a>

### StakeContractConfig
StakeContractConfig initial setup config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_bond` | [uint64](#uint64) |  |  |
| `tokens_per_point` | [uint64](#uint64) |  |  |
| `unbonding_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `claim_autoreturn_limit` | [uint32](#uint32) |  |  |






<a name="mage-coven.poe.v1beta1.TG4Member"></a>

### TG4Member
TG4Member member of the Engagement group.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address is the bech32 address string |
| `points` | [uint64](#uint64) |  |  |






<a name="mage-coven.poe.v1beta1.ValidatorVotingContractConfig"></a>

### ValidatorVotingContractConfig
ValidatorVotingContractConfig CommunityPoolContractConfig


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_rules` | [VotingRules](#mage-coven.poe.v1beta1.VotingRules) |  | VotingRules rules for the tally |






<a name="mage-coven.poe.v1beta1.ValsetContractConfig"></a>

### ValsetContractConfig
ValsetContractConfig initial setup config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_points` | [uint64](#uint64) |  |  |
| `max_validators` | [uint32](#uint32) |  | MaxValidators The maximum number of validators that can be included in the Tendermint validator set. |
| `epoch_length` | [google.protobuf.Duration](#google.protobuf.Duration) |  | EpochLength We update the Tendermint validator set only once per epoch. Epoch # is env.block.time/epoch_length (round down). The first block with a new epoch number will trigger a new validator calculation. |
| `epoch_reward` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `scaling` | [uint32](#uint32) |  | Scaling is the factor to multiply cw4-group weights to produce the Tendermint validator power |
| `fee_percentage` | [string](#string) |  | FeePercentage is the percentage of total accumulated fees that is subtracted from tokens minted as rewards. 50% by default. To disable this feature just set it to 0 (which effectively means that fees don't affect the per-epoch reward). |
| `community_pool_reward_ratio` | [string](#string) |  | CommunityPoolRewardRation in percentage |
| `engagement_reward_ratio` | [string](#string) |  | EngagementRewardRatio reward ration in percentage for all |
| `validator_reward_ratio` | [string](#string) |  | ValidatorRewardRation in percentage for all |
| `auto_unjail` | [bool](#bool) |  | AutoUnjail if set to true, we will auto-unjail any validator after their jailtime is over. |
| `double_sign_slash_ratio` | [string](#string) |  | DoubleSignSlashRatio Validators who are caught double signing are jailed forever and their bonded tokens are slashed based on this value. |
| `verify_validators` | [bool](#bool) |  | When a validator joins the valset, verify they sign the first block since joining or jail them for a period otherwise.

The verification happens every time the validator becomes an active validator, including when they are unjailed or when they just gain enough power to participate. |
| `offlineJailDuration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | OfflineJailDuration duration to jail a validator for in case they don't sign their first epoch boundary block. After the period, they have to pass verification again, ad infinitum. |






<a name="mage-coven.poe.v1beta1.VotingRules"></a>

### VotingRules
VotingRules contains configuration for the tally.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_period` | [uint32](#uint32) |  | VotingPeriod Voting period in days |
| `quorum` | [string](#string) |  | Quorum voting quorum percentage (1-100) |
| `threshold` | [string](#string) |  | Threshold voting threshold percentage (50-100) |
| `allow_end_early` | [bool](#bool) |  | AllowEndEarly If true, and absolute threshold and quorum are met, we can end before voting period finished. (Recommended value: true, unless you have special needs) |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/poe/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/poe/v1beta1/query.proto



<a name="mage-coven.poe.v1beta1.QueryContractAddressRequest"></a>

### QueryContractAddressRequest
QueryContractAddressRequest is the request type for the Query/ContractAddress
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_type` | [PoEContractType](#mage-coven.poe.v1beta1.PoEContractType) |  | ContractType is the type of contract |






<a name="mage-coven.poe.v1beta1.QueryContractAddressResponse"></a>

### QueryContractAddressResponse
QueryContractAddressRequest is the response type for the
Query/ContractAddress RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="mage-coven.poe.v1beta1.QueryUnbondingPeriodRequest"></a>

### QueryUnbondingPeriodRequest
QueryUnbondingPeriodRequest is request type for the Query/UnbondingPeriod RPC
method






<a name="mage-coven.poe.v1beta1.QueryUnbondingPeriodResponse"></a>

### QueryUnbondingPeriodResponse
QueryUnbondingPeriodResponse is response type for the Query/UnbondingPeriod
RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `time` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Time is the time that must pass |






<a name="mage-coven.poe.v1beta1.QueryValidatorDelegationRequest"></a>

### QueryValidatorDelegationRequest
QueryValidatorDelegationRequest is request type for the
Query/ValidatorDelegation RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="mage-coven.poe.v1beta1.QueryValidatorDelegationResponse"></a>

### QueryValidatorDelegationResponse
QueryValidatorDelegationResponse is response type for the
Query/ValidatorDelegation RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="mage-coven.poe.v1beta1.QueryValidatorEngagementRewardRequest"></a>

### QueryValidatorEngagementRewardRequest
QueryValidatorEngagementRewardRequest is the request type for the
Query/ValidatorEngagementReward RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |






<a name="mage-coven.poe.v1beta1.QueryValidatorEngagementRewardResponse"></a>

### QueryValidatorEngagementRewardResponse
QueryValidatorEngagementRewardResponse is the response type for the
Query/ValidatorEngagementReward RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  |  |






<a name="mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardRequest"></a>

### QueryValidatorOutstandingRewardRequest
QueryValidatorOutstandingRewardRequest is the request type for the
Query/ValidatorOutstandingReward RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |






<a name="mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardResponse"></a>

### QueryValidatorOutstandingRewardResponse
QueryValidatorOutstandingRewardResponse is the response type for the
Query/ValidatorOutstandingReward RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  |  |






<a name="mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsRequest"></a>

### QueryValidatorUnbondingDelegationsRequest
QueryValidatorUnbondingDelegationsRequest is required type for the
Query/ValidatorUnbondingDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsResponse"></a>

### QueryValidatorUnbondingDelegationsResponse
QueryValidatorUnbondingDelegationsResponse is response type for the
Query/ValidatorUnbondingDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `entries` | [cosmos.staking.v1beta1.UnbondingDelegationEntry](#cosmos.staking.v1beta1.UnbondingDelegationEntry) | repeated | unbonding delegation entries |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="mage-coven.poe.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ContractAddress` | [QueryContractAddressRequest](#mage-coven.poe.v1beta1.QueryContractAddressRequest) | [QueryContractAddressResponse](#mage-coven.poe.v1beta1.QueryContractAddressResponse) | ContractAddress queries the address for one of the PoE contracts | GET|/gridiron/poe/v1beta1/contract/{contract_type}|
| `Validators` | [.cosmos.staking.v1beta1.QueryValidatorsRequest](#cosmos.staking.v1beta1.QueryValidatorsRequest) | [.cosmos.staking.v1beta1.QueryValidatorsResponse](#cosmos.staking.v1beta1.QueryValidatorsResponse) | Validators queries all validators that match the given status. | GET|/gridiron/poe/v1beta1/validators|
| `Validator` | [.cosmos.staking.v1beta1.QueryValidatorRequest](#cosmos.staking.v1beta1.QueryValidatorRequest) | [.cosmos.staking.v1beta1.QueryValidatorResponse](#cosmos.staking.v1beta1.QueryValidatorResponse) | Validator queries validator info for given validator address. | GET|/gridiron/poe/v1beta1/validators/{validator_addr}|
| `UnbondingPeriod` | [QueryUnbondingPeriodRequest](#mage-coven.poe.v1beta1.QueryUnbondingPeriodRequest) | [QueryUnbondingPeriodResponse](#mage-coven.poe.v1beta1.QueryUnbondingPeriodResponse) | Validator queries validator info for given validator address. | GET|/gridiron/poe/v1beta1/unbonding|
| `ValidatorDelegation` | [QueryValidatorDelegationRequest](#mage-coven.poe.v1beta1.QueryValidatorDelegationRequest) | [QueryValidatorDelegationResponse](#mage-coven.poe.v1beta1.QueryValidatorDelegationResponse) | ValidatorDelegation queries self delegated amount for given validator. | GET|/poe/gridiron/v1beta1/validators/{validator_addr}/delegation|
| `ValidatorUnbondingDelegations` | [QueryValidatorUnbondingDelegationsRequest](#mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsRequest) | [QueryValidatorUnbondingDelegationsResponse](#mage-coven.poe.v1beta1.QueryValidatorUnbondingDelegationsResponse) | ValidatorUnbondingDelegations queries unbonding delegations of a validator. | GET|/gridiron/poe/v1beta1/validators/{validator_addr}/unbonding_delegations|
| `HistoricalInfo` | [.cosmos.staking.v1beta1.QueryHistoricalInfoRequest](#cosmos.staking.v1beta1.QueryHistoricalInfoRequest) | [.cosmos.staking.v1beta1.QueryHistoricalInfoResponse](#cosmos.staking.v1beta1.QueryHistoricalInfoResponse) | HistoricalInfo queries the historical info for given height. | GET|/gridiron/poe/v1beta1/historical_info/{height}|
| `ValidatorOutstandingReward` | [QueryValidatorOutstandingRewardRequest](#mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardRequest) | [QueryValidatorOutstandingRewardResponse](#mage-coven.poe.v1beta1.QueryValidatorOutstandingRewardResponse) | ValidatorOutstandingRewards queries rewards of a validator address. | GET|/gridiron/poe/v1beta1/validators/{validator_address}/outstanding_reward|
| `ValidatorEngagementReward` | [QueryValidatorEngagementRewardRequest](#mage-coven.poe.v1beta1.QueryValidatorEngagementRewardRequest) | [QueryValidatorEngagementRewardResponse](#mage-coven.poe.v1beta1.QueryValidatorEngagementRewardResponse) | ValidatorEngagementReward queries rewards of a validator address. | GET|/gridiron/poe/v1beta1/validators/{validator_address}/engagement_reward|

 <!-- end services -->



<a name="mage-coven/poe/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/poe/v1beta1/tx.proto



<a name="mage-coven.poe.v1beta1.MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator defines a PoE message for creating a new validator.
Based on the SDK staking.MsgCreateValidator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [cosmos.staking.v1beta1.Description](#cosmos.staking.v1beta1.Description) |  | Description meta data |
| `operator_address` | [string](#string) |  | OperatorAddress is the bech32 address string |
| `pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  | Pubkey public key |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Amount defines the initial staking amount from a liquid balance |
| `vesting_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | VestingAmount defines the initial staking amount from a vesting account balance |






<a name="mage-coven.poe.v1beta1.MsgCreateValidatorResponse"></a>

### MsgCreateValidatorResponse
MsgCreateValidatorResponse defines the MsgCreateValidator response type.






<a name="mage-coven.poe.v1beta1.MsgDelegate"></a>

### MsgDelegate
MsgDelegate defines a SDK message for performing a self delegation of coins
by a node operator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `vesting_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="mage-coven.poe.v1beta1.MsgDelegateResponse"></a>

### MsgDelegateResponse
MsgDelegateResponse defines the Msg/Delegate response type.






<a name="mage-coven.poe.v1beta1.MsgUndelegate"></a>

### MsgUndelegate
MsgUndelegate defines a SDK message for performing an undelegation from a
node operator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="mage-coven.poe.v1beta1.MsgUndelegateResponse"></a>

### MsgUndelegateResponse
MsgUndelegateResponse defines the Msg/Undelegate response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="mage-coven.poe.v1beta1.MsgUpdateValidator"></a>

### MsgUpdateValidator
MsgCreateValidator defines a PoE message for updating validator metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [cosmos.staking.v1beta1.Description](#cosmos.staking.v1beta1.Description) |  | New Description meta data |
| `operator_address` | [string](#string) |  | OperatorAddress is the bech32 address string Also know as "signer" in other messages |






<a name="mage-coven.poe.v1beta1.MsgUpdateValidatorResponse"></a>

### MsgUpdateValidatorResponse
MsgUpdateValidatorResponse defines the MsgUpdateValidator response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="mage-coven.poe.v1beta1.Msg"></a>

### Msg
Msg defines the staking Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateValidator` | [MsgCreateValidator](#mage-coven.poe.v1beta1.MsgCreateValidator) | [MsgCreateValidatorResponse](#mage-coven.poe.v1beta1.MsgCreateValidatorResponse) | CreateValidator defines a method for creating a new validator. | |
| `UpdateValidator` | [MsgUpdateValidator](#mage-coven.poe.v1beta1.MsgUpdateValidator) | [MsgUpdateValidatorResponse](#mage-coven.poe.v1beta1.MsgUpdateValidatorResponse) | MsgCreateValidator defines a method for updating validator metadata | |
| `Delegate` | [MsgDelegate](#mage-coven.poe.v1beta1.MsgDelegate) | [MsgDelegateResponse](#mage-coven.poe.v1beta1.MsgDelegateResponse) | Delegate defines a method for performing a self delegation of coins by a node operator | |
| `Undelegate` | [MsgUndelegate](#mage-coven.poe.v1beta1.MsgUndelegate) | [MsgUndelegateResponse](#mage-coven.poe.v1beta1.MsgUndelegateResponse) | Undelegate defines a method for performing an undelegation from a node operator | |

 <!-- end services -->



<a name="mage-coven/twasm/v1beta1/contract_extension.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/twasm/v1beta1/contract_extension.proto



<a name="mage-coven.twasm.v1beta1.RegisteredPrivilege"></a>

### RegisteredPrivilege
RegisteredPrivilege stores position and privilege name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `position` | [uint32](#uint32) |  |  |
| `privilege_type` | [string](#string) |  |  |






<a name="mage-coven.twasm.v1beta1.GridironContractDetails"></a>

### GridironContractDetails
GridironContractDetails is a custom extension to the wasmd ContractInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `registered_privileges` | [RegisteredPrivilege](#mage-coven.twasm.v1beta1.RegisteredPrivilege) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/twasm/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/twasm/v1beta1/genesis.proto



<a name="mage-coven.twasm.v1beta1.Contract"></a>

### Contract
Contract struct encompasses ContractAddress, ContractInfo, and ContractState


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `contract_info` | [cosmwasm.wasm.v1.ContractInfo](#cosmwasm.wasm.v1.ContractInfo) |  |  |
| `kv_model` | [KVModel](#mage-coven.twasm.v1beta1.KVModel) |  |  |
| `custom_model` | [CustomModel](#mage-coven.twasm.v1beta1.CustomModel) |  |  |






<a name="mage-coven.twasm.v1beta1.CustomModel"></a>

### CustomModel
CustomModel contains the raw json data for a contract to seed its state on
import


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg` | [bytes](#bytes) |  | Msg json encoded message to be passed to the contract on import |






<a name="mage-coven.twasm.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [cosmwasm.wasm.v1.Params](#cosmwasm.wasm.v1.Params) |  | Params sdk type Params for wasmd |
| `codes` | [cosmwasm.wasm.v1.Code](#cosmwasm.wasm.v1.Code) | repeated | Codes has all stored wasm codes and metadata |
| `contracts` | [Contract](#mage-coven.twasm.v1beta1.Contract) | repeated | Contracts contains all instantiated contracts, state and metadata |
| `sequences` | [cosmwasm.wasm.v1.Sequence](#cosmwasm.wasm.v1.Sequence) | repeated | Sequences names and values |
| `gen_msgs` | [cosmwasm.wasm.v1.GenesisState.GenMsgs](#cosmwasm.wasm.v1.GenesisState.GenMsgs) | repeated | GenMsgs has wasmd sdk type messages to execute in the genesis phase |
| `privileged_contract_addresses` | [string](#string) | repeated | PrivilegedContractAddresses is a list of contract addresses that can have special permissions |
| `pinned_code_ids` | [uint64](#uint64) | repeated | PinnedCodeIDs has codeInfo ids for wasm codes that are pinned in cache |






<a name="mage-coven.twasm.v1beta1.KVModel"></a>

### KVModel
KVModel is a wrapper around the wasmd default key value model.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `models` | [cosmwasm.wasm.v1.Model](#cosmwasm.wasm.v1.Model) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/twasm/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/twasm/v1beta1/proposal.proto



<a name="mage-coven.twasm.v1beta1.DemotePrivilegedContractProposal"></a>

### DemotePrivilegedContractProposal
PromoteToPrivilegedContractProposal gov proposal content type to remove
"privileges" from a contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |






<a name="mage-coven.twasm.v1beta1.PromoteToPrivilegedContractProposal"></a>

### PromoteToPrivilegedContractProposal
PromoteToPrivilegedContractProposal gov proposal content type to add
"privileges" to a contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | Title is a short summary |
| `description` | [string](#string) |  | Description is a human readable text |
| `contract` | [string](#string) |  | Contract is the address of the smart contract |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="mage-coven/twasm/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mage-coven/twasm/v1beta1/query.proto



<a name="mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeRequest"></a>

### QueryContractsByPrivilegeTypeRequest
QueryContractsByPrivilegeTypeRequest is the request type for the
Query/ContractsByPrivilegeType RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `privilege_type` | [string](#string) |  |  |






<a name="mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeResponse"></a>

### QueryContractsByPrivilegeTypeResponse
QueryContractsByPrivilegeTypeResponse is the response type for the
Query/ContractsByPrivilegeType RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [string](#string) | repeated | contracts are a set of contract addresses |






<a name="mage-coven.twasm.v1beta1.QueryPrivilegedContractsRequest"></a>

### QueryPrivilegedContractsRequest
QueryPrivilegedContractsResponse is the request type for the
Query/PrivilegedContracts RPC method






<a name="mage-coven.twasm.v1beta1.QueryPrivilegedContractsResponse"></a>

### QueryPrivilegedContractsResponse
QueryPrivilegedContractsResponse is the response type for the
Query/PrivilegedContracts RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contracts` | [string](#string) | repeated | contracts are a set of contract addresses |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="mage-coven.twasm.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PrivilegedContracts` | [QueryPrivilegedContractsRequest](#mage-coven.twasm.v1beta1.QueryPrivilegedContractsRequest) | [QueryPrivilegedContractsResponse](#mage-coven.twasm.v1beta1.QueryPrivilegedContractsResponse) | PrivilegedContracts returns all privileged contracts | GET|/gridiron/twasm/v1beta1/contracts/privileged|
| `ContractsByPrivilegeType` | [QueryContractsByPrivilegeTypeRequest](#mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeRequest) | [QueryContractsByPrivilegeTypeResponse](#mage-coven.twasm.v1beta1.QueryContractsByPrivilegeTypeResponse) | ContractsByPrivilegeType returns all contracts that have registered for the privilege type | GET|/gridiron/twasm/v1beta1/contracts/privilege/{privilege_type}|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

