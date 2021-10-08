package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/budget/x/budget/types"
)

// DONTCOVER

// Simulation parameter constants
const (
	EpochBlocks = "epoch_blocks"
	Budgets     = "budgets"
)

// GenEpochBlocks returns randomized epoch blocks.
func GenEpochBlocks(r *rand.Rand) uint32 {
	return uint32(simulation.RandIntBetween(r, int(types.DefaultEpochBlocks), 10))
}

// GenGenBudgets returns randomized budgets.
func GenBudgets(r *rand.Rand) []types.Budget {
	ranBudgets := make([]types.Budget, 0)
	// TODO:
	// if budget source address is the same, then we should lower rate and a number of budgets.
	// Otherwise, get an error "the total rate of budgets with the same budget source address value should not exceed 1"
	//
	// 1. consider to use randomized BudgetSourceAddress
	// 2. use Cosmos Hub's FeeCollector module account and reduce rate and number of budgets
	//
	for i := 0; i < simulation.RandIntBetween(r, 1, 2); i++ {
		budget := types.Budget{
			Name:                "simulation-test-" + simulation.RandStringOfLength(r, 5),
			Rate:                sdk.NewDecFromIntWithPrec(sdk.NewInt(int64(simulation.RandIntBetween(r, 1, 2))), 1),
			BudgetSourceAddress: "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta", // Cosmos Hub's FeeCollector module account
			CollectionAddress:   sdk.AccAddress(address.Module(types.ModuleName, []byte("GravityDEXFarmingBudget"))).String(),
			StartTime:           types.ParseTime("2000-01-01T00:00:00Z"),
			EndTime:             types.ParseTime("9999-12-31T00:00:00Z"),
		}
		ranBudgets = append(ranBudgets, budget)
	}

	return ranBudgets
}

// RandomizedGenState generates a random GenesisState for budget.
func RandomizedGenState(simState *module.SimulationState) {
	var epochBlocks uint32
	var budgets []types.Budget
	simState.AppParams.GetOrGenerate(
		simState.Cdc, EpochBlocks, &epochBlocks, simState.Rand,
		func(r *rand.Rand) { epochBlocks = GenEpochBlocks(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, Budgets, &budgets, simState.Rand,
		func(r *rand.Rand) { budgets = GenBudgets(r) },
	)

	budgetGenesis := types.GenesisState{
		Params: types.Params{
			EpochBlocks: epochBlocks,
			Budgets:     budgets,
		},
	}

	bz, _ := json.MarshalIndent(&budgetGenesis, "", " ")
	fmt.Printf("Selected randomly generated budget parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&budgetGenesis)
}
