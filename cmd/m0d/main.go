package main

import (
	"os"
	"strings"

	"github.com/liubaninc/m0/app"
	"github.com/liubaninc/m0/cmd/m0d/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, strings.ToUpper(app.Name), app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
