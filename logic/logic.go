package logic

import (
	"bytes"
	"embed"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// TODO: Add option to add .direnv to .gitignore

//go:embed data/*
var content embed.FS

const (
	flakeDirectory = "data"
	flakeFilename  = "flake.nix"
	envrc          = ".envrc"
)

func CreateFlakeFile(name string) error {
	selectedFlake := name

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(pwd, flakeFilename)); err == nil {
		log.Fatal("Error creating flake.nix:", "File already exists")
	}

	file, err := os.Create(filepath.Join(pwd, flakeFilename))
	if err != nil {
		return err
	}

	flakeFile, err := content.ReadFile(filepath.Join(flakeDirectory, selectedFlake, "flake.nix"))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, bytes.NewReader(flakeFile))
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func CreateEnvRCInCWD() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(cwd, envrc))
	if err != nil {
		return err
	}

	_, err = file.WriteString("use flake")
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

func InitGitRepository() error {
	cmd := exec.Command("git", "init")
	if cmd.Err != nil {
		return cmd.Err
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "add", ".envrc", "flake.nix")
	if cmd.Err != nil {
		return cmd.Err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func DirNamesFromEmbededDir() []string {
	dir, err := content.ReadDir(flakeDirectory)
	if err != nil {
		log.Fatal("Error reading Embedded FS:", err)
	}

	items := make([]string, len(dir))

	for i, v := range dir {
		items[i] = v.Name()
	}

	return items
}
