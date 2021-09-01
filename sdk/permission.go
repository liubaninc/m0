package sdk

import (
	"context"
	permissiontypes "github.com/liubaninc/m0/x/permission/types"
)

func (c Client) GetPerms(address string) ([]string, error) {
	queryClient := permissiontypes.NewQueryClient(c)
	res, err := queryClient.Account(context.Background(), &permissiontypes.QueryGetAccountRequest{
		Address: address,
	})
	if err != nil {
		return nil, err
	}
	return res.Account.Perms, nil
}
