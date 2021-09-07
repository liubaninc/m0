package main

import (
	"flag"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gorm.io/gorm"
	"os"

	"github.com/liubaninc/m0/cmd/synced/api"
	_ "github.com/liubaninc/m0/cmd/synced/docs"
	"github.com/liubaninc/m0/cmd/synced/model"
	"github.com/liubaninc/m0/cmd/synced/syncer"
	msdk "github.com/liubaninc/m0/sdk"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server synced server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.Parse()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	//Keyring.Key("faucet")
	kr, err := keyring.New("m0", keyring.BackendMemory, ".", os.Stdin)
	if err != nil {
		panic(err)
	}
	if _, err := kr.NewAccount("faucet", *faucetAccount, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1); err != nil {
		panic(err)
	}
	client := msdk.MustNew(*RPCHost, kr)

	var conn gorm.Dialector
	switch *mode {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName)
		conn = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", *dbHost, *dbUser, *dbPassword, *dbName, *dbPort)
		conn = postgres.Open(dsn)
	case "sqlserver":

	case "sqlite":
		conn = sqlite.Open(*dbName + ".db")
	default:
		panic("not support db mode")
	}
	model := model.New(conn, logger)
	syncer.New(model.DB, client, logger).Run()
	api.New(model.DB, client, logger).Run(*port)
}

var RPCHost = flag.String("rpc-host", flagRPCHost, "node rpc host")
var mode = flag.String("db-mode", flagDBMode, "database mode, mysql,postgres,sqlserver,sqlite")
var dbHost = flag.String("db-host", flagDBHost, "database host")
var dbPort = flag.Int("db-port", flagDBPort, "database port")
var dbUser = flag.String("db-user", flagDBUser, "database user")
var dbPassword = flag.String("db-pass", flagDBPassword, "database password")
var dbName = flag.String("db-name", flagDBName, "database name")
var port = flag.Int("port", flagPort, "listen port")
var faucetAccount = flag.String("faucet-account", flagFaucetAccount, "faucet account mnemonic")
var faucetCoin = flag.String("faucet-coin", flagFaucetCoin, "faucet amount")

const (
	flagRPCHost       = "tcp://localhost:26657"
	flagDBMode        = "mysql"
	flagDBHost        = "localhost"
	flagDBPort        = 3306
	flagDBName        = "m0"
	flagDBUser        = "root"
	flagDBPassword    = "123456"
	flagPort          = 8080
	flagFaucetAccount = "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
	flagFaucetCoin    = "10m0token"
)
