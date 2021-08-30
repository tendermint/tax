package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Querier) Budgets(c context.Context, req *types.QueryBudgetsRequest) (*types.QueryBudgetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	var budgets []types.BudgetResponse
	for _, b := range params.Budgets {
		// TODO: filter budgets
		collectedCoins := k.GetTotalCollectedCoins(ctx, b.Name)
		budgets = append(budgets, types.BudgetResponse{
			Budget:              b,
			TotalCollectedCoins: collectedCoins,
		})
	}

	return &types.QueryBudgetsResponse{Budgets: budgets}, nil
}