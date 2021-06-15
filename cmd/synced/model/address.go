package model

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Address struct {
	ID            uint      `json:"-" gorm:"primarykey" `
	Name          string    `json:"name" gorm:"unique"`
	Address       string    `json:"address" gorm:"unique"`
	AccountNumber uint64    `json:"account_number"`
	Sequence      uint64    `json:"sequence"`
	Balance       string    `json:"-"`
	Coins         sdk.Coins `json:"coins" gorm:"-" `
}

func (addr *Address) Fill(names ...string) {
	coins, _ := sdk.ParseCoinsNormalized(strings.TrimRight(addr.Balance, ","))
	if len(names) == 0 {
		addr.Coins = coins
	} else {
		addr.Coins = sdk.NewCoins()
		for _, name := range names {
			addr.Coins = addr.Coins.Add(sdk.NewCoin(name, coins.AmountOf(name)))
		}
	}
}
