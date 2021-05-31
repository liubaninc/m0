package keeper

import (
	"errors"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewUtxoReader(ctx sdk.Context, k Keeper) *UtxoReader {
	return &UtxoReader{
		ctx: ctx,
		k:   k,
	}
}

// UtxoReader utxos interface for reader
type UtxoReader struct {
	k   Keeper
	ctx sdk.Context
}

func (u *UtxoReader) SelectUtxos(fromAddr sdk.AccAddress, totalNeed sdk.Coins, lock int64) ([]*utxotypes.Input, sdk.Coins, error) {
	return u.k.utxoKeeper.SelectUtxos(u.ctx, fromAddr, totalNeed, lock)
}

// UtxoCache 合约内资产转移缓存
type UtxoCache struct {
	utxovm      *UtxoReader
	inputCache  []*utxotypes.Input
	outputCache []*utxotypes.Output
	inputUsed   map[string]bool
	isPenetrate bool
}

func NewUtxoCache(utxovm *UtxoReader) *UtxoCache {
	return &UtxoCache{
		utxovm:      utxovm,
		inputUsed:   make(map[string]bool),
		isPenetrate: true,
	}
}

func NewUtxoCacheWithInputs(inputs []*utxotypes.Input) *UtxoCache {
	return &UtxoCache{
		inputCache:  inputs,
		inputUsed:   make(map[string]bool),
		isPenetrate: false,
	}
}

func (u *UtxoCache) selectUtxos(from sdk.AccAddress, amount sdk.Coin) (sdk.Int, error) {
	if u.isPenetrate {
		inputs, total, err := u.utxovm.SelectUtxos(from, sdk.NewCoins(amount), 60)
		if err != nil {
			return sdk.Int{}, err
		}
		u.inputCache = append(u.inputCache, inputs...)
		return total.AmountOf(amount.Denom), nil
	}

	sum := sdk.ZeroInt()
	for _, input := range u.inputCache {
		// Since contract calls bridge serially, a mismatched from address is an error
		if input.FromAddr != from.String() {
			return sdk.Int{}, errors.New("from address mismatch in utxo cache")
		}
		if input.Amount.Denom != amount.Denom {
			continue
		}
		if _, ok := u.inputUsed[input.Index()]; ok {
			continue
		}
		u.inputUsed[input.Index()] = true
		sum = sum.Add(input.Amount.Amount)
		if sum.GTE(amount.Amount) {
			break
		}
	}
	if sum.LT(amount.Amount) {
		return sdk.Int{}, errors.New("utxo not enough in utxo cache")
	}
	return sum, nil
}

func (u *UtxoCache) Transfer(from, to sdk.AccAddress, amount sdk.Coins) error {
	if amount.IsZero() {
		return nil
	}
	for _, coin := range amount {
		total, err := u.selectUtxos(from, coin)
		if err != nil {
			return err
		}
		u.outputCache = append(u.outputCache, &utxotypes.Output{
			ToAddr:      to.String(),
			Amount:       coin,
			FrozenHeight: 0,
		})
		// 找零
		if sub := total.Sub(coin.Amount); sub.IsPositive() {
			u.outputCache = append(u.outputCache, &utxotypes.Output{
				ToAddr:      from.String(),
				Amount:       sdk.NewCoin(coin.Denom, sub),
				FrozenHeight: 0,
			})
		}
	}
	return nil
}

func (u *UtxoCache) GetRWSets() ([]*utxotypes.Input, []*utxotypes.Output) {
	if u.isPenetrate {
		return u.inputCache, u.outputCache
	}
	var inputs []*utxotypes.Input
	for _, input := range u.inputCache {
		if _, ok := u.inputUsed[input.Index()]; ok {
			inputs = append(inputs, input)
		}
	}
	return inputs, u.outputCache
}
