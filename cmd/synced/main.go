package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gorm.io/gorm"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/liubaninc/m0/cmd/synced/api"
	_ "github.com/liubaninc/m0/cmd/synced/docs"
	"github.com/liubaninc/m0/cmd/synced/model"
	"github.com/liubaninc/m0/cmd/synced/syncer"
	msdk "github.com/liubaninc/m0/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
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
	rootCmd := &cobra.Command{
		Use:   "synced",
		Short: "m0 synced Daemon (server)",
	}

	rootCmd.AddCommand(newStartCommand())
	rootCmd.AddCommand(server.VersionCmd())
	executor := cli.PrepareMainCmd(rootCmd, "SYNCED", os.ExpandEnv("$HOME/.synced"))
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

const (
	flagRPCHost       = "node-rpc"
	flagDBMode        = "db-mode"
	flagDBHost        = "db-host"
	flagDBPort        = "db-port"
	flagDBName        = "db-name"
	flagDBUser        = "db-user"
	flagDBPassword    = "db-pass"
	flagPort          = "port"
	flagFaucetAccount = "faucet-account"
	flagFaucetCoin    = "faucet-amount"
)

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start synced service",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

			//Keyring.Key("faucet")
			kr, err := keyring.New("m0", keyring.BackendMemory, viper.GetString(flags.FlagHome), os.Stdin)
			if err != nil {
				panic(err)
			}
			if _, err := kr.NewAccount("faucet", viper.GetString(flagFaucetAccount), keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1); err != nil {
				panic(err)
			}
			mode := viper.GetString(flagDBMode)
			client := msdk.MustNew(viper.GetString(flagRPCHost), kr)
			dbUser := viper.GetString(flagDBUser)
			dbPassword := viper.GetString(flagDBPassword)
			dbHost := viper.GetString(flagDBHost)
			dbPort := viper.GetInt(flagDBPort)
			dbName := viper.GetString(flagDBName)
			var conn gorm.Dialector
			switch mode {
			case "mysql":
				dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
				conn = mysql.Open(dsn)
			case "postgres":
				dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbName, dbPort)
				conn = postgres.Open(dsn)
			case "sqlserver":

			case "sqlite":
				conn = sqlite.Open(dbName + ".db")
			default:
				panic("not support db mode")
			}
			model := model.New(conn, logger)
			syncer.New(model.DB, client, logger).Run()
			api.New(model.DB, client, logger).Run(viper.GetInt(flagPort))
			return nil
		},
	}
	cmd.Flags().String(flagRPCHost, "tcp://localhost:26657", "node rpc host")
	cmd.Flags().String(flagDBMode, "sqlite", "database mode, mysql,postgres,sqlserver,sqlite")
	cmd.Flags().String(flagDBHost, "localhost", "database host")
	cmd.Flags().Int(flagDBPort, 3306, "database port")
	cmd.Flags().String(flagDBUser, "root", "database user")
	cmd.Flags().String(flagDBPassword, "root", "database password")
	cmd.Flags().String(flagDBName, "m0", "database name")
	cmd.Flags().Int(flagPort, 8080, "listen port")
	cmd.Flags().String(flagFaucetAccount, "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear",
		"faucet account mnemonic")
	cmd.Flags().String(flagFaucetCoin, "10m0token", "faucet amount")
	return cmd
}
