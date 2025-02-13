package main

import (
	"github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.AddCommand(cli.InitStoresCmd())
	log.Fatal(rootCmd.Execute())
}
