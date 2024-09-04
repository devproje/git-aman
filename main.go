package main

import (
	"flag"
	"fmt"
	"github.com/devproje/git-aman/action"
	"github.com/devproje/git-aman/profile"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
	"os/exec"
)

var (
	list   bool
	create bool
	profId int
)

const VERSION = "1.0.0-alpha.1"

func init() {
	flag.BoolVar(&list, "l", false, "list all profiles")
	flag.BoolVar(&create, "c", false, "create profile")
	flag.IntVar(&profId, "p", 0, "specify profile to use")
	flag.Parse()

	profile.LoadAll()
}

func main() {
	if create && list {
		return
	}

	log.Printf("git-aman %s", VERSION)
	log.Printf("DATA DIRECTORY: %s", util.GetDataDir())

	err := checkGit()
	if err != nil {
		log.Fatalln(err)
	}

	if list {
		profile.QueryProfs()
		return
	}

	if create {
		action.ProfileCreate()
		return
	}

	if profId == 0 {
		log.Errorln("profile id must not be null")
		return
	}

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

	_, err = os.Open(fmt.Sprintf("%s/.git-credentials", util.GetHome()))
	if err != nil {
		// TODO: create new one
		log.Fatalln(err)
		return
	}

	updateGitConfig(prof.Name, prof.Email)

	uri := fmt.Sprintf("%s://%s:%s@%s",
		prof.AuthData.Protocol,
		prof.AuthData.Username,
		prof.AuthData.Secret,
		prof.AuthData.Server)
	err = os.WriteFile(fmt.Sprintf("%s/.git-credentials", util.GetHome()), []byte(uri), 200)
	if err != nil {
		return
	}
}

func updateGitConfig(name, email string) {
	_, err := exec.Command("git", "config", "--global", "user.name", name).Output()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = exec.Command("git", "config", "--global", "user.email", email).Output()
	if err != nil {
		log.Fatalln(err)
	}
}
