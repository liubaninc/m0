package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/spf13/cobra"
)

var (
	commandFuncs []func() *cobra.Command

	rootOptions RootOptions
)

var (
	logger = log.New(ioutil.Discard, "xdev ", log.LstdFlags|log.Lshortfile)
)

type RootOptions struct {
}

func addCommand(cmdFunc func() *cobra.Command) {
	commandFuncs = append(commandFuncs, cmdFunc)
}

func rootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "xdev",
		SilenceErrors: false,
		SilenceUsage:  true,
	}
	initRootCommand(rootCmd)
	for _, cmdFunc := range commandFuncs {
		rootCmd.AddCommand(cmdFunc())
	}
	rootCmd.AddCommand(server.VersionCmd())
	return rootCmd
}

func initRootCommand(cmd *cobra.Command) {
	var verbose bool
	rootFlags := cmd.PersistentFlags()
	rootFlags.BoolVarP(&verbose, "verbose", "v", false, "show debug message")
	cobra.OnInitialize(func() {
		if verbose {
			logger.SetOutput(os.Stderr)
		}
	})
}

func Main() {
	root := rootCommand()
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
