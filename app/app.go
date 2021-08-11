package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	appparams "github.com/liubaninc/m0/app/params"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport
	mibcmodule "github.com/liubaninc/m0/x/mibc"
	mibcmodulekeeper "github.com/liubaninc/m0/x/mibc/keeper"
	mibcmoduletypes "github.com/liubaninc/m0/x/mibc/types"
	peermodule "github.com/liubaninc/m0/x/peer"
	peermodulekeeper "github.com/liubaninc/m0/x/peer/keeper"
	peermoduletypes "github.com/liubaninc/m0/x/peer/types"
	permissionmodule "github.com/liubaninc/m0/x/permission"
	permissionmodulekeeper "github.com/liubaninc/m0/x/permission/keeper"
	permissionmoduletypes "github.com/liubaninc/m0/x/permission/types"
	pkimodule "github.com/liubaninc/m0/x/pki"
	pkimodulekeeper "github.com/liubaninc/m0/x/pki/keeper"
	pkimoduletypes "github.com/liubaninc/m0/x/pki/types"
	"github.com/liubaninc/m0/x/utxo"
	utxokeeper "github.com/liubaninc/m0/x/utxo/keeper"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	validatormodule "github.com/liubaninc/m0/x/validator"
	validatormodulekeeper "github.com/liubaninc/m0/x/validator/keeper"
	validatormoduletypes "github.com/liubaninc/m0/x/validator/types"
	"github.com/liubaninc/m0/x/wasm"
	wasmkeeper "github.com/liubaninc/m0/x/wasm/keeper"
	wasmtypes "github.com/liubaninc/m0/x/wasm/types"
)

const Name = "m0"

// this line is used by starport scaffolding # stargate/wasm/app/enabledProposals

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		params.AppModuleBasic{},
		ibc.AppModuleBasic{},
		// this line is used by starport scaffolding # stargate/app/moduleBasic
		peermodule.AppModuleBasic{},
		mibcmodule.AppModuleBasic{},
		pkimodule.AppModuleBasic{},
		permissionmodule.AppModuleBasic{},
		validatormodule.AppModuleBasic{},
		wasm.AppModuleBasic{},
		utxo.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName: nil,
		utxotypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		mibcmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner},
	}
)

var (
	_ CosmosApp               = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Marshaler
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly

	// make scoped keepers public for test purposes
	ScopedIBCKeeper capabilitykeeper.ScopedKeeper

	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	PeerKeeper       peermodulekeeper.Keeper
	ScopedMibcKeeper capabilitykeeper.ScopedKeeper
	MibcKeeper       mibcmodulekeeper.Keeper

	PkiKeeper pkimodulekeeper.Keeper

	PermissionKeeper permissionmodulekeeper.Keeper

	ValidatorKeeper validatormodulekeeper.Keeper

	UtxoKeeper utxokeeper.Keeper

	WasmKeeper wasmkeeper.Keeper

	// the module manager
	mm *module.Manager
}

// New returns a reference to an initialized Gaia.
// NewSimApp returns a reference to an initialized SimApp.
func New(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig appparams.EncodingConfig,
	// this line is used by starport scaffolding # stargate/app/newArgument
	appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *App {

	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
		peermoduletypes.StoreKey,
		mibcmoduletypes.StoreKey,
		pkimoduletypes.StoreKey,
		permissionmoduletypes.StoreKey,
		validatormoduletypes.StoreKey,
		utxotypes.StoreKey,
		wasmtypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &App{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	// this line is used by starport scaffolding # stargate/app/scopedKeeper

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
	)
	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.ValidatorKeeper, scopedIBCKeeper,
	)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition
	app.UtxoKeeper = *utxokeeper.NewKeeper(
		appCodec,
		keys[utxotypes.StoreKey],
		keys[utxotypes.MemStoreKey],
		app.GetSubspace(utxotypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
	)
	utxoModule := utxo.NewAppModule(appCodec, app.UtxoKeeper)

	scopedMibcKeeper := app.CapabilityKeeper.ScopeToModule(mibcmoduletypes.ModuleName)
	app.ScopedMibcKeeper = scopedMibcKeeper
	app.MibcKeeper = *mibcmodulekeeper.NewKeeper(
		appCodec,
		keys[mibcmoduletypes.StoreKey],
		keys[mibcmoduletypes.MemStoreKey],
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedMibcKeeper,
		app.UtxoKeeper,
	)
	mibcModule := mibcmodule.NewAppModule(appCodec, app.MibcKeeper)

	app.PkiKeeper = *pkimodulekeeper.NewKeeper(
		appCodec,
		keys[pkimoduletypes.StoreKey],
		keys[pkimoduletypes.MemStoreKey],
	)
	pkiModule := pkimodule.NewAppModule(appCodec, app.PkiKeeper)

	app.PeerKeeper = *peermodulekeeper.NewKeeper(
		appCodec,
		keys[peermoduletypes.StoreKey],
		keys[peermoduletypes.MemStoreKey],
		app.GetSubspace(peermoduletypes.ModuleName),
		app.PkiKeeper,
	)
	peerModule := peermodule.NewAppModule(appCodec, app.PeerKeeper)

	app.PermissionKeeper = *permissionmodulekeeper.NewKeeper(
		appCodec,
		keys[permissionmoduletypes.StoreKey],
		keys[permissionmoduletypes.MemStoreKey],
		app.GetSubspace(permissionmoduletypes.ModuleName),
		app.AccountKeeper,
	)
	permissionModule := permissionmodule.NewAppModule(appCodec, app.PermissionKeeper)

	app.ValidatorKeeper = *validatormodulekeeper.NewKeeper(
		appCodec,
		keys[validatormoduletypes.StoreKey],
		keys[validatormoduletypes.MemStoreKey],
		app.GetSubspace(validatormoduletypes.ModuleName),
	)
	validatorModule := validatormodule.NewAppModule(appCodec, app.ValidatorKeeper)

	app.WasmKeeper = *wasmkeeper.NewKeeper(
		appCodec,
		keys[wasmtypes.StoreKey],
		keys[wasmtypes.MemStoreKey],
		app.GetSubspace(wasmtypes.ModuleName),
		app.UtxoKeeper,
	)
	wasmModule := wasm.NewAppModule(appCodec, app.WasmKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	// this line is used by starport scaffolding # ibc/app/router
	ibcRouter.AddRoute(mibcmoduletypes.ModuleName, mibcModule)
	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.ValidatorKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		// this line is used by starport scaffolding # stargate/app/appModule
		peerModule,
		mibcModule,
		pkiModule,
		permissionModule,
		validatorModule,
		utxoModule,
		wasmModule,
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		ibchost.ModuleName,
		validatormoduletypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(validatormoduletypes.ModuleName)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		ibchost.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
		peermoduletypes.ModuleName,
		mibcmoduletypes.ModuleName,
		pkimoduletypes.ModuleName,
		permissionmoduletypes.ModuleName,
		validatormoduletypes.ModuleName,
		wasmtypes.ModuleName,
		utxotypes.ModuleName,
		//
		genutiltypes.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		sdk.ChainAnteDecorators(
			ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
			ante.NewRejectExtensionOptionsDecorator(),
			ante.NewMempoolFeeDecorator(),
			ante.NewValidateBasicDecorator(),
			ante.TxTimeoutHeightDecorator{},
			ante.NewValidateMemoDecorator(app.AccountKeeper),
			ante.NewConsumeGasForTxSizeDecorator(app.AccountKeeper),
			ante.NewRejectFeeGranterDecorator(),
			ante.NewSetPubKeyDecorator(app.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
			ante.NewValidateSigCountDecorator(app.AccountKeeper),
			// ante.NewDeductFeeDecorator(app.AccountKeeper, app.BankKeeper),
			ante.NewSigGasConsumeDecorator(app.AccountKeeper, ante.DefaultSigVerificationGasConsumer),
			ante.NewSigVerificationDecorator(app.AccountKeeper, encodingConfig.TxConfig.SignModeHandler()),
			ante.NewIncrementSequenceDecorator(app.AccountKeeper),
			ante.NewValidatePermissonDecorator(app.PermissionKeeper),
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	app.SetIDPeerFilter(func(info string) abci.ResponseQuery {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("SetIDPeerFilter", r)
				debug.PrintStack()
			}
		}()
		return app.PeerKeeper.IDPeerFilter(app.GetPeerFilterContext(), info)
	})

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		app.CapabilityKeeper.InitializeAndSeal(ctx)
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	// this line is used by starport scaffolding # stargate/app/beforeInitReturn

	return app
}

// Name returns the name of the App
func (app *App) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *App) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *App) AppCodec() codec.Marshaler {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry
func (app *App) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *App) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *App) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *App) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryMarshaler, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace
	paramsKeeper.Subspace(peermoduletypes.ModuleName)
	paramsKeeper.Subspace(mibcmoduletypes.ModuleName)
	paramsKeeper.Subspace(pkimoduletypes.ModuleName)
	paramsKeeper.Subspace(permissionmoduletypes.ModuleName)
	paramsKeeper.Subspace(validatormoduletypes.ModuleName)
	paramsKeeper.Subspace(wasmtypes.ModuleName)
	paramsKeeper.Subspace(utxotypes.ModuleName)

	return paramsKeeper
}
