package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgDestroy{}, "meng/Destroy", nil)

	cdc.RegisterConcrete(&MsgUnfreeze{}, "meng/Unfreeze", nil)

	cdc.RegisterConcrete(&MsgFreeze{}, "meng/Freeze", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDestroy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFreeze{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
