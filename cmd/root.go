package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

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
		list, err := cmd.PersistentFlags().GetBool("list")
		if err != nil {
			log.Fatalln(err)
		}

		if list {
			var out strings.Builder

			for _, v := range logic.DirNamesFromEmbededDir() {
				_, err = out.WriteString(fmt.Sprintf("%s\n", v))
				if err != nil {
					log.Fatalln(err)
				}
			}
			fmt.Printf("Available shells:\n%s", out.String())
			return
		}

		if len(args) == 0 {
			ui.Run()
		} else {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}

			err = logic.CopyFilesFromEmbededDir(args[0], cwd)
			if err != nil {
				log.Fatalln(err)
			}

			err = logic.CreateEnvRC(cwd)
			if err != nil {
				log.Fatalln(err)
			}

			git, err := cmd.PersistentFlags().GetBool("git")
			if err != nil {
				log.Fatalln(err)
			}

			if git {
				err = logic.InitGitRepository()
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	},
	ValidArgs: logic.DirNamesFromEmbededDir(),
}

func init() {
	rootCmd.PersistentFlags().BoolP("git", "g", false, "Initialize Git repository")
	rootCmd.PersistentFlags().BoolP("list", "l", false, "List available shells")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
