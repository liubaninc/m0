package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgUnfreezeCert{}, "pki/UnfreezeCert", nil)

	cdc.RegisterConcrete(&MsgFreezeCert{}, "pki/FreezeCert", nil)

	cdc.RegisterConcrete(&MsgRevokeRootCert{}, "pki/RevokeRootCert", nil)

	cdc.RegisterConcrete(&MsgRevokeCert{}, "pki/RevokeCert", nil)

	cdc.RegisterConcrete(&MsgAddCert{}, "pki/AddCert", nil)

	cdc.RegisterConcrete(&MsgAddRootCert{}, "pki/AddRootCert", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreezeCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFreezeCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeRootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRootCert{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
