package cmd

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"path/filepath"

	"github.com/cosmos/go-bip39"
	cfg "github.com/tendermint/tendermint/config"
	tmed25519 "github.com/tendermint/tendermint/crypto/ed25519"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// InitializeNodeValidatorFiles creates private validator and p2p configuration files using the given mnemonic.
// If no valid mnemonic is given, a random one will be used instead.
func InitializeNodeValidatorFilesFromMnemonic(config *cfg.Config, mnemonic string) (nodeID string, valPubKey cryptotypes.PubKey, err error) {
	if len(mnemonic) > 0 && !bip39.IsMnemonicValid(mnemonic) {
		return "", nil, fmt.Errorf("invalid mnemonic")
	}

	nodeKeyFile := config.NodeKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(nodeKeyFile), 0777); err != nil {
		return "", nil, err
	}

	nodeKey := &p2p.NodeKey{
		PrivKey: tmed25519.GenPrivKeyFromSecret([]byte(mnemonic)),
	}

	if err := nodeKey.SaveAs(nodeKeyFile); err != nil {
		return "", nil, err
	}

	nodeID = string(nodeKey.ID())

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return "", nil, err
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return "", nil, err
	}

	var filePV *privval.FilePV
	if len(mnemonic) == 0 {
		filePV = privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	} else {
		privKey := tmed25519.GenPrivKeyFromSecret([]byte(mnemonic))
		filePV = privval.NewFilePV(privKey, pvKeyFile, pvStateFile)
		filePV.Save()
	}

	tmValPubKey, err := filePV.GetPubKey()
	if err != nil {
		return "", nil, err
	}

	valPubKey, err = cryptocodec.FromTmPubKeyInterface(tmValPubKey)
	if err != nil {
		return "", nil, err
	}

	return nodeID, valPubKey, nil
}

func SaveCoinKey(keybase keyring.Keyring, keyName string, mnemonic string, overwrite bool, algo keyring.SignatureAlgo) (sdk.AccAddress, error) {
	exists := false
	_, err := keybase.Key(keyName)
	if err == nil {
		exists = true
	}

	// ensure no overwrite
	if !overwrite && exists {
		return sdk.AccAddress([]byte{}), fmt.Errorf(
			"key already exists, overwrite is disabled")
	}

	// generate a private key, with recovery phrase
	if exists {
		err = keybase.Delete(keyName)
		if err != nil {
			return sdk.AccAddress([]byte{}), fmt.Errorf(
				"failed to overwrite key")
		}
	}

	info, err := keybase.NewAccount(keyName, mnemonic ,keyring.DefaultBIP39Passphrase, sdk.FullFundraiserPath, algo)
	if err != nil {
		return sdk.AccAddress([]byte{}), err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), nil
}