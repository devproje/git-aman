package main

import (
	"flag"
	"fmt"
	"github.com/devproje/git-aman/action"
	"github.com/devproje/git-aman/config"
	"github.com/devproje/git-aman/profile"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
)

var (
	list   bool
	create bool
	profId int
)

const VERSION = "1.0.0-alpha.2"

func init() {
	flag.BoolVar(&list, "l", false, "list all profiles")
	flag.BoolVar(&create, "c", false, "create profile")
	flag.IntVar(&profId, "p", 0, "specify profile to use")
	flag.Parse()

	config.Conf = config.Load()
	profile.Load()
}

func main() {
	if create && list {
		return
	}

	log.Printf("git-aman %s", VERSION)
	log.Printf("DATA DIRECTORY: %s", util.GetDataDir())

	err := util.CheckGit()
	if err != nil {
		log.Fatalln("git package not installed")
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

func change(prof *profile.Profile) {
	if prof == nil {
		log.Fatalln(fmt.Errorf("profile %d is null", profId))
		return
	}

	secretPath := fmt.Sprintf("%s/.git-credentials", util.GetHome())
	_, err := os.Open(secretPath)
	if err != nil {
		if _, err = os.Create(secretPath); err != nil {
			log.Fatalln(err)
			return
		}
	}

	util.UpdateGitConfig(prof.Name, prof.Email)
	uri := fmt.Sprintf("%s://%s:%s@%s",
		prof.AuthData.Protocol,
		prof.AuthData.Username,
		prof.AuthData.Secret,
		prof.AuthData.Server)
	err = os.WriteFile(fmt.Sprintf("%s/.git-credentials", util.GetHome()), []byte(uri), 200)
	if err != nil {
		return
	}

	config.SetId(prof.Id)
}
