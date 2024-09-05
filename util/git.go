package util

import (
	"github.com/devproje/plog/log"
	"os/exec"
)

func CheckGit() error {
	_, err := exec.Command("git", "--version").Output()
	if err != nil {
		return err
	}

	return nil
}

func UpdateGitConfig(name, email string) {
	_, err := exec.Command("git", "config", "--global", "user.name", name).Output()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = exec.Command("git", "config", "--global", "user.email", email).Output()
	if err != nil {
		log.Fatalln(err)
	}
}
