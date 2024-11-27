package logic

import (
	"embed"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed data/*
var content embed.FS

const (
	flakeDirectory = "data"
	direnv         = ".direnv"
	envrc          = ".envrc"
)

func CopyFilesFromEmbededDir(dirName string, dstPath string) error {
	srcDirectoryPath := filepath.Join(flakeDirectory, dirName)

	files, err := content.ReadDir(srcDirectoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		dstFilePath := filepath.Join(dstPath, file.Name())

		srcFile, err := content.Open(filepath.Join(srcDirectoryPath, file.Name()))
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstFilePath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateEnvRC(path string) error {
	file, err := os.Create(filepath.Join(path, envrc))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("use flake")
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

	err = addFileToGitIgnore(direnv)
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

func addFileToGitIgnore(name string) error {
	file, err := os.OpenFile(".gitignore", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(name)
	if err != nil {
		return err
	}

	return nil
}
