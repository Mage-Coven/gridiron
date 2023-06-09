package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ slashingtypes.QueryServer = &LegacySlashingGRPCQuerier{}

type LegacySlashingGRPCQuerier struct {
	keeper ContractSource
}

func NewLegacySlashingGRPCQuerier(keeper ContractSource) *LegacySlashingGRPCQuerier { //nolint:golint
	return &LegacySlashingGRPCQuerier{keeper: keeper}
}

// SigningInfo legacy support for cosmos-sdk signing info. Note that not all field are available on gridiron
func (g LegacySlashingGRPCQuerier) SigningInfo(c context.Context, req *slashingtypes.QuerySigningInfoRequest) (*slashingtypes.QuerySigningInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	valAddr, err := sdk.AccAddressFromBech32(req.ConsAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "validator address")
	}
	return &slashingtypes.QuerySigningInfoResponse{
		ValSigningInfo: slashingtypes.ValidatorSigningInfo{
			Address:             valAddr.String(),
			StartHeight:         0,
			IndexOffset:         0,
			Tombstoned:          false,
			MissedBlocksCounter: 0,
		},
	}, nil
}

// SigningInfos is not supported and will return unimplemented error
func (g LegacySlashingGRPCQuerier) SigningInfos(c context.Context, req *slashingtypes.QuerySigningInfosRequest) (*slashingtypes.QuerySigningInfosResponse, error) {
	logNotImplemented(sdk.UnwrapSDKContext(c), "SigningInfos")
	return nil, status.Error(codes.Unimplemented, "not available, yet")
}

// Params is not supported. Method returns default slashing module params.
func (g LegacySlashingGRPCQuerier) Params(c context.Context, req *slashingtypes.QueryParamsRequest) (*slashingtypes.QueryParamsResponse, error) {
	return &slashingtypes.QueryParamsResponse{
		Params: slashingtypes.Params{
			SignedBlocksWindow:      0,
			MinSignedPerWindow:      sdk.ZeroDec(),
			DowntimeJailDuration:    0,
			SlashFractionDoubleSign: sdk.ZeroDec(),
			SlashFractionDowntime:   sdk.ZeroDec(),
		},
	}, nil
}
