package kernel

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
)

const (
	ContractNumberBucket      = "ContractNumber"
	Contract2AccountBucket    = "Contract2Account"
	Account2ContractBucket    = "Account2Contract"
	Account2ContractSeparator = "\x01"
)

// contractMethods manage methods about contract
type contractMethods struct {
	xbridge *bridge.XBridge
}

// Deploy deploys contract
func (c *contractMethods) Deploy(ctx *KContext, args map[string][]byte) (*contract.Response, error) {
	// check if account exist
	accountName := []byte(ctx.Initiator)
	contractName := args["contract_name"]
	if accountName == nil || contractName == nil {
		return nil, errors.New("invoke DeployMethod error, account name or contract name is nil")
	}
	// check if contractName is ok
	if contractErr := ValidContractName(string(contractName)); contractErr != nil {
		return nil, fmt.Errorf("deploy failed, contract `%s` contains illegal character, error: %s", contractName, contractErr)
	}

	number := uint64(0)
	//value, _ := ctx.ModelCache.Get(ContractNumberBucket, []byte("number"))
	//if value == nil {
	//	number = 0
	//} else {
	//	number, _ = strconv.ParseUint(string(value.PureData.Value), 10, 64)
	//}

	out, resourceUsed, err := c.xbridge.DeployContract(ctx.ContextConfig, args, number)
	if err != nil {
		return nil, err
	}
	ctx.AddResourceUsed(resourceUsed)

	// key: contract, value: account
	err = ctx.ModelCache.Put(Contract2AccountBucket, contractName, accountName)
	if err != nil {
		return nil, err
	}
	err = ctx.ModelCache.Put(Account2ContractBucket, []byte(string(accountName)+Account2ContractSeparator+string(contractName)), []byte("true"))
	if err != nil {
		return nil, err
	}
	//err = ctx.ModelCache.Put(ContractNumberBucket, []byte("number"), []byte(strconv.FormatUint(number+1, 10)))
	//if err != nil {
	//	return nil, err
	//}
	return out, nil
}

// Upgrade upgrades contract
func (c *contractMethods) Upgrade(ctx *KContext, args map[string][]byte) (*contract.Response, error) {
	contractName := args["contract_name"]
	if contractName == nil {
		return nil, errors.New("invoke Upgrade error, contract name is nil")
	}

	ver, err := ctx.ModelCache.Get(Contract2AccountBucket, contractName)
	if err != nil {
		return nil, err
	}
	if ctx.Initiator != string(ver.GetPureData().Value) {
		return nil, errors.New("verify contract owner permission failed")
	}

	resp, resourceUsed, err := c.xbridge.UpgradeContract(ctx.ContextConfig, args)
	if err != nil {
		return nil, err
	}
	ctx.AddResourceUsed(resourceUsed)
	return resp, nil
}

var (
	contractNameRegex = regexp.MustCompile("^[a-zA-Z_]{1}[0-9a-zA-Z_.]+[0-9a-zA-Z_]$")
)

func ValidContractName(contractName string) error {
	if !contractNameRegex.MatchString(contractName) {
		return fmt.Errorf("contract name does not fit the rule %v of contract name", contractNameRegex.String())
	}
	return nil
}
