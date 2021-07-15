package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgApproveRevokeAccountRequest{}, "authority/ApproveRevokeAccount", nil)

	cdc.RegisterConcrete(&MsgProposeRevokeAccountRequest{}, "authority/ProposeRevokeAccount", nil)

	cdc.RegisterConcrete(&MsgApproveAddAccountRequest{}, "authority/ApproveAddAccount", nil)

	cdc.RegisterConcrete(&MsgProposeAddAccountRequest{}, "authority/ProposeAddAccount", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeAccountRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeAccountRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddAccountRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddAccountRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
