package pki

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/keeper"
	"github.com/liubaninc/m0/x/pki/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the certificate
	for _, elem := range genState.CertificateList {
		k.SetCertificate(ctx, *elem)
	}

	// Set all the certificates
	for _, elem := range genState.CertificatesList {
		k.SetCertificates(ctx, *elem)
	}

	for _, elem := range genState.RevokeCertificatesList {
		k.SetRevokedCertificates(ctx, *elem)
	}

	for _, elem := range genState.ChildCertificatesList {
		k.SetChildCertificates(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all certificate
	certificateList := k.GetAllCertificate(ctx)
	for _, elem := range certificateList {
		elem := elem
		genesis.CertificateList = append(genesis.CertificateList, &elem)
	}

	// Get all certificates
	certificatesList := k.GetAllCertificates(ctx)
	for _, elem := range certificatesList {
		elem := elem
		genesis.CertificatesList = append(genesis.CertificatesList, &elem)
	}

	revokeCertificatesList := k.GetAllRevokedCertificates(ctx)
	for _, elem := range revokeCertificatesList {
		elem := elem
		genesis.CertificatesList = append(genesis.CertificatesList, &elem)
	}

	childCertificatesList := k.GetAllChildCertificates(ctx)
	for _, elem := range childCertificatesList {
		elem := elem
		genesis.CertificatesList = append(genesis.CertificatesList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
