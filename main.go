package main

import (
	"flag"
	"fmt"
	"github.com/devproje/git-aman/profile"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
	"os/exec"
)

var (
	list   bool
	profId int
)

const VERSION = "1.0.0-alpha.1"

func init() {
	flag.BoolVar(&list, "l", false, "list all profiles")
	flag.IntVar(&profId, "p", 0, "specify profile to use")
	flag.Parse()

	profile.LoadAll()
}

func main() {
	if list {
		profile.QueryProfs()
		return
	}

	log.Printf("git-aman %s", VERSION)
	log.Printf("DATA DIRECTORY: %s", util.GetDataDir())

	change(profile.Read(profId))
}

func checkGit() error {
	_, err := exec.Command("git", "--version").Output()
	if err != nil {
		return err
	}

	return nil
}

func change(prof *profile.Profile) {
	err := checkGit()
	if err != nil {
		log.Panicln("git not installed\n")
	}

	if prof == nil {
		log.Fatalln(fmt.Errorf("profile %d is null", profId))
		return
	}

	config, err := os.Open("gitconfig")
	if err != nil {
		// TODO: create new one
		log.Fatalln(err)
		return
	}
	defer config.Close()

	credential, err := os.Open(".git-credentials")
	if err != nil {
		// TODO: create new one
		log.Fatalln(err)
		return
	}
	defer credential.Close()

	// TODO: create change config

	uri := fmt.Sprintf("https://%s:%s@%s", prof.AuthData.Username, prof.AuthData.Secret, prof.AuthData.Server)
	_, err = credential.Write([]byte(uri))
	if err != nil {
		return
	}
}
