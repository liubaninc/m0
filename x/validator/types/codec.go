package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgLeaveValidator{}, "validator/LeaveValidator", nil)

	cdc.RegisterConcrete(&MsgEditValidator{}, "validator/EditValidator", nil)

	cdc.RegisterConcrete(&MsgCreateValidator{}, "validator/CreateValidator", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLeaveValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditValidator{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
