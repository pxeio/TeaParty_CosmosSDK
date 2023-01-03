package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitSell{}, "party/SubmitSell", nil)
	cdc.RegisterConcrete(&MsgBuy{}, "party/Buy", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "party/Cancel", nil)
	cdc.RegisterConcrete(&MsgAccountWatchOutcome{}, "party/AccountWatchOutcome", nil)
	cdc.RegisterConcrete(&MsgAccountWatchFailure{}, "party/AccountWatchFailure", nil)
	cdc.RegisterConcrete(&MsgCreateOrdersAwaitingFinalizer{}, "party/CreateOrdersAwaitingFinalizer", nil)
	cdc.RegisterConcrete(&MsgUpdateOrdersAwaitingFinalizer{}, "party/UpdateOrdersAwaitingFinalizer", nil)
	cdc.RegisterConcrete(&MsgDeleteOrdersAwaitingFinalizer{}, "party/DeleteOrdersAwaitingFinalizer", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitSell{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccountWatchOutcome{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAccountWatchFailure{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrdersAwaitingFinalizer{},
		&MsgUpdateOrdersAwaitingFinalizer{},
		&MsgDeleteOrdersAwaitingFinalizer{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
