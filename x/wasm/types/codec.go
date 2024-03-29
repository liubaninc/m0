package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgApproveDeployContract{}, "wasm/ApproveDeployContract", nil)

	cdc.RegisterConcrete(&MsgProposeDeployContract{}, "wasm/ProposeDeployContract", nil)

	cdc.RegisterConcrete(&MsgUndeploy{}, "wasm/Undeploy", nil)

	cdc.RegisterConcrete(&MsgUnfreeze{}, "wasm/Unfreeze", nil)

	cdc.RegisterConcrete(&MsgFreeze{}, "wasm/Freeze", nil)

	cdc.RegisterConcrete(&MsgDeploy{}, "wasm/Deploy", nil)
	cdc.RegisterConcrete(&MsgInvoke{}, "wasm/Invoke", nil)
	cdc.RegisterConcrete(&MsgUpgrade{}, "wasm/Upgrade", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveDeployContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeDeployContract{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndeploy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeploy{}, &MsgInvoke{}, &MsgUpgrade{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
