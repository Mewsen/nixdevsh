package cmd

import (
	"log"

	"github.com/mewsen/nixdevsh/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixdevsh",
	Short: "nixdevsh is a nix dev-shell Generator",
	Long:  `A fast Nix development shell Generator built with Go.`,
	Example: `'nixdevsh rust' initiate rust env
'nixdevsh go --git' initiate go env with git repository`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			ui.StartUI()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
