package cmd

import (
	"log"

	"github.com/mewsen/nixdevsh/logic"
	"github.com/mewsen/nixdevsh/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixdevsh",
	Short: "nixdevsh is a nix dev-shell Generator",
	Long:  `A fast Nix development shell Generator built with Go.`,
	Example: `'nixdevsh rust' initiate rust env
'nixdevsh go --git' initiate go env with git repository`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return err
		}

		err = cobra.RangeArgs(0, 1)(cmd, args)
		if err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ui.Run()
		} else {

			logic.CreateFlakeFile(args[0])
			logic.CreateEnvRCInCWD()

			git, err := cmd.PersistentFlags().GetBool("git")
			if err != nil {
				log.Fatalln(err)
			}

			if git {
				logic.InitGitRepository()
			}
		}
	},
	ValidArgs: logic.DirNamesFromEmbededDir(),
}

func init() {
	rootCmd.PersistentFlags().BoolP("git", "g", false, "Initialize Git repository")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
