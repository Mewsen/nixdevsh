package ui

import (
	"log"

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

	{
		err = logic.CreateFlakeFile(ui.shellSelection)
		if err != nil {
			log.Fatalln(err)
		}

		err = logic.CreateEnvRCInCWD()
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
}

func Run() {
	ui := Ui{}

	ui.Run()
}
