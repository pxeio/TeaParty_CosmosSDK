package types_test

import (
	"testing"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				TradeOrdersList: []types.TradeOrders{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				PendingOrdersList: []types.PendingOrders{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				OrdersAwaitingFinalizerList: []types.OrdersAwaitingFinalizer{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated tradeOrders",
			genState: &types.GenesisState{
				TradeOrdersList: []types.TradeOrders{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pendingOrders",
			genState: &types.GenesisState{
				PendingOrdersList: []types.PendingOrders{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated ordersAwaitingFinalizer",
			genState: &types.GenesisState{
				OrdersAwaitingFinalizerList: []types.OrdersAwaitingFinalizer{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
