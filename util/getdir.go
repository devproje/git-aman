package util

import (
	"fmt"
	"github.com/devproje/plog/log"
	"os"
)

func GetHome() string {
	return os.Getenv("HOME")
}

func GetDataDir() string {
	path := fmt.Sprintf("%s/.cache/git-aman", GetHome())
	if _, err := os.ReadFile(path); err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalln(fmt.Errorf("can not create directory: %s", path))
		}
	}

	return path
}
