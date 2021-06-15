package main

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"os"

	"github.com/liubaninc/m0/cmd/synced/api"
	_ "github.com/liubaninc/m0/cmd/synced/docs"
	"github.com/liubaninc/m0/cmd/synced/model"
	"github.com/liubaninc/m0/cmd/synced/syncer"
	msdk "github.com/liubaninc/m0/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/server"
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

// @host 8.131.229.225:8080
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
	executor := cli.PrepareMainCmd(rootCmd, "M0_SYNCED", os.ExpandEnv("$HOME/.m0synced"))
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

const (
	flagRPCHost    = "node-rpc"
	flagDBHost     = "db-host"
	flagDBPort     = "db-port"
	flagDBName     = "db-name"
	flagDBUser     = "db-user"
	flagDBPassword = "db-pass"
	flagPort       = "port"
)

func newStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start synced service",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

			kr, _ := keyring.New("m0", keyring.BackendMemory, viper.GetString(flags.FlagHome), os.Stdin)
			client := msdk.MustNew(viper.GetString(flagRPCHost), kr)
			model := model.New(viper.GetString(flagDBHost), viper.GetInt(flagDBPort), viper.GetString(flagDBUser), viper.GetString(flagDBPassword), viper.GetString(flagDBName), logger)
			syncer.New(model.DB, client, logger).Run()
			api.New(model.DB, client, logger).Run(viper.GetInt(flagPort))
			return nil
		},
	}
	cmd.Flags().String(flagRPCHost, "tcp://localhost:26657", "node rpc host")
	cmd.Flags().String(flagDBHost, "localhost", "database host")
	cmd.Flags().Int(flagDBPort, 3306, "database port")
	cmd.Flags().String(flagDBUser, "root", "database user")
	cmd.Flags().String(flagDBPassword, "root", "database password")
	cmd.Flags().String(flagDBName, "m0", "database name")
	cmd.Flags().Int(flagPort, 8080, "listen port")
	return cmd
}
