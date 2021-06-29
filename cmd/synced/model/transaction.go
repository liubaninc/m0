package model

import (
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Transaction struct {
	ID           uint       `json:"-" gorm:"primarykey" `
	Hash         string     `json:"hash" gorm:"unique" `
	Size         int        `json:"size"`
	Memo         string     `json:"memo" `
	Fee          string     `json:"fee" `
	Status       bool       `json:"status" `
	Height       int64      `json:"height" gorm:"index" `
	Time         string     `json:"time" `
	MsgNum       int        `json:"msg_num" `
	Type         string     `json:"type"`
	UTXOMsgs     []*MsgUTXO `json:"utxo_msgs,omitempty"`
	Confirmed    int64      `json:"confirmed" `
	Version      string     `json:"version" gorm:"-" `
	Assets       string     `json:"assets" `
	Contracts    string     `json:"-" `
	Addresses    string     `json:"-" `
	Address      string     `json:"address,omitempty" gorm:"-"`
	AddressCoins sdk.Coins  `json:"coins,omitempty" gorm:"-" `
}

func (t *Transaction) FillConfirmed(height int64, address string, coins ...string) {
	t.Confirmed = height - t.Height + 1
	if len(address) != 0 {
		mapCoins := map[string]bool{}
		for _, coin := range coins {
			if len(coin) == 0 {
				continue
			}
			mapCoins[coin] = true
		}

		inputs := sdk.NewCoins()
		outputs := sdk.NewCoins()
		for _, msg := range t.UTXOMsgs {
			for _, input := range msg.Inputs {
				if strings.Compare(address, input.Address) == 0 {
					coin, _ := sdk.ParseCoinNormalized(input.Amount)
					if len(mapCoins) > 0 && !mapCoins[coin.Denom] {
						continue
					}
					inputs = inputs.Add(coin)
				}
			}
			for _, output := range msg.Outputs {
				if strings.Compare(address, output.Address) == 0 {
					coin, _ := sdk.ParseCoinNormalized(output.Amount)
					if len(mapCoins) > 0 && !mapCoins[coin.Denom] {
						continue
					}
					outputs = outputs.Add(coin)
				}
			}
		}
		t.Address = address
		t.AddressCoins, _ = outputs.SafeSub(inputs)
	} else {
		inputs := sdk.NewCoins()
		for _, msg := range t.UTXOMsgs {
			for _, input := range msg.Inputs {
				coin, _ := sdk.ParseCoinNormalized(input.Amount)
				inputs = inputs.Add(coin)
			}
		}
		t.AddressCoins = inputs
	}
	t.Assets = strings.Trim(t.Assets, ",")
}

func FillConfirmed(txs []*Transaction, height int64, address string, coins ...string) []*Transaction {
	var ttxs []*Transaction
	for _, tx := range txs {
		tx.FillConfirmed(height, address, coins...)
		ttxs = append(ttxs, tx)
	}
	return ttxs
}

type MTransaction struct {
	ID           uint       `json:"-" gorm:"primarykey" `
	Hash         string     `json:"hash" gorm:"unique" `
	Size         int        `json:"size"`
	Memo         string     `json:"memo" `
	Fee          string     `json:"fee" `
	Status       bool       `json:"status" `
	Height       int64      `json:"height" gorm:"index" `
	Time         string     `json:"time" `
	MsgNum       int        `json:"msg_num" `
	Type         string     `json:"type" `
	UTXOMsgs     []*MsgUTXO `json:"utxo_msgs,omitempty" gorm:"-"`
	Confirmed    int64      `json:"confirmed" `
	Version      string     `json:"version" gorm:"-" `
	Assets       string     `json:"-" `
	Contracts    string     `json:"-" `
	Addresses    string     `json:"-" `
	Address      string     `json:"address,omitempty" gorm:"-"`
	AddressCoins sdk.Coins  `json:"coins" gorm:"-" `

	// added
	Msgs       string   `json:"-"`
	Raw        string   `json:"-" `
	Signature  string   `json:"-"`                   // 已签名数
	Signatures []string `json:"signatures" gorm:"-"` // 已签名数
}

func (t *MTransaction) FillConfirmed(address string, coins ...string) {
	// added
	if len(t.Signature) > 0 {
		t.Signatures = strings.Split(t.Signature, ",")
	} else {
		t.Signatures = []string{}
	}
	t.UTXOMsgs = []*MsgUTXO{}

	json.Unmarshal([]byte(t.Msgs), &t.UTXOMsgs)

	if len(address) != 0 {
		mapCoins := map[string]bool{}
		for _, coin := range coins {
			if len(coin) == 0 {
				continue
			}
			mapCoins[coin] = true
		}

		inputs := sdk.NewCoins()
		outputs := sdk.NewCoins()
		for _, msg := range t.UTXOMsgs {
			for _, input := range msg.Inputs {
				if strings.Compare(address, input.Address) == 0 {
					coin, _ := sdk.ParseCoinNormalized(input.Amount)
					if len(mapCoins) > 0 && !mapCoins[coin.Denom] {
						continue
					}
					inputs = inputs.Add(coin)
				}
			}
			for _, output := range msg.Outputs {
				if strings.Compare(address, output.Address) == 0 {
					coin, _ := sdk.ParseCoinNormalized(output.Amount)
					if len(mapCoins) > 0 && !mapCoins[coin.Denom] {
						continue
					}
					outputs = outputs.Add(coin)
				}
			}
		}
		t.Address = address
		t.AddressCoins, _ = outputs.SafeSub(inputs)
	} else {
		inputs := sdk.NewCoins()
		for _, msg := range t.UTXOMsgs {
			for _, input := range msg.Inputs {
				coin, _ := sdk.ParseCoinNormalized(input.Amount)
				inputs = inputs.Add(coin)
			}
		}
		t.AddressCoins = inputs
	}
	t.Assets = strings.Trim(t.Assets, ",")
}
