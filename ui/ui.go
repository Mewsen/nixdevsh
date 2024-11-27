package ui

import (
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/mewsen/nixdevsh/logic"
)

type Ui struct {
	shellSelection string
	initGitRepo    bool
}

func (ui Ui) Run() {
	shells := logic.DirNamesFromEmbededDir()

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a development shell").
				Options(huh.NewOptions(shells...)...).
				Value(&ui.shellSelection),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Initialize Git repository?").
				Value(&ui.initGitRepo),
		),
	).Run()
	if err != nil {
		log.Fatalln(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	err = logic.CopyFilesFromEmbededDir(ui.shellSelection, cwd)
	if err != nil {
		log.Fatalln(err)
	}

	err = logic.CreateEnvRC(cwd)
	if err != nil {
		log.Fatalln(err)
	}

	if ui.initGitRepo {
		err = logic.InitGitRepository()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Run() {
	ui := Ui{}

	ui.Run()
}
