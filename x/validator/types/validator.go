package types

import (
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const DoNotModifyDesc = "[do-not-modify]"
const DoNotModifyPower = int64(-1)

// ABCI ValidatorUpdate message to add new validator to validator set.
func (v *Validator) ABCIValidatorUpdate() abci.ValidatorUpdate {
	pub, err := cryptocodec.ToTmProtoPublicKey(sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.PubKey))
	if err != nil {
		panic(err)
	}
	return abci.ValidatorUpdate{
		PubKey: pub,
		Power:  v.GetPower(),
	}
}

// ABCI ValidatorUpdate message to remove validator from validator set.
func (v *Validator) ABCIValidatorUpdateZero() abci.ValidatorUpdate {
	pub, err := cryptocodec.ToTmProtoPublicKey(sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.PubKey))
	if err != nil {
		panic(err)
	}
	return abci.ValidatorUpdate{
		PubKey: pub,
		Power:  0,
	}
}

// Accessors
func (v *Validator) GetOperator() sdk.ValAddress {
	addr, err := sdk.ValAddressFromBech32(v.Creator)
	if err != nil {
		panic(err)
	}
	return addr
}

func (v *Validator) GetConsPubKey() types.PubKey {
	return sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.PubKey)
}
func (v *Validator) GetConsAddr() sdk.ConsAddress {
	return sdk.ConsAddress(v.GetConsPubKey().Address())
}

func (v *Validator) ValidateBasic() error {
	if v.GetOperator().Empty() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	if v.GetPubKey() == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "missing consensus pubkey")
	}
	return nil
}

func (d Description) ValidateBasic() error {
	if len(d.Moniker) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid Description Name: it cannot be empty")
	}
	return nil
}
