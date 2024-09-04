package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/devproje/git-aman/profile"
	"github.com/devproje/git-aman/types"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"golang.org/x/term"
	"os"
	"os/exec"
	"strings"
	"syscall"
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

	if list {
		profile.QueryProfs()
		return
	}

	if create {
		createProf()
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

func createProf() {
	var displayName []byte
	var user, email []byte
	var protocol types.Protocol
	var server, username []byte

	var reader = bufio.NewReader(os.Stdin)

	fmt.Printf("type your profile display name: ")
	displayName, _ = reader.ReadBytes('\n')

	fmt.Printf("type git config 'user.name': ")
	user, _ = reader.ReadBytes('\n')

	fmt.Printf("type git config 'user.email': ")
	email, _ = reader.ReadBytes('\n')

	fmt.Printf("type your git protocol: ")
	raw, _ := reader.ReadBytes('\n')
	protocol = types.CheckProto(string(raw))

	fmt.Printf("type your git server: ")
	server, _ = reader.ReadBytes('\n')

	fmt.Printf("type your git username: ")
	username, _ = reader.ReadBytes('\n')

	fmt.Printf("type your git password: ")
	secret, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	data := profile.Profile{
		Id:          profile.ProfSize() + 1,
		DisplayName: strings.ReplaceAll(string(displayName), "\n", ""),
		Name:        strings.ReplaceAll(string(user), "\n", ""),
		Email:       strings.ReplaceAll(string(email), "\n", ""),
		AuthData: profile.Credential{
			Protocol: protocol,
			Server:   strings.ReplaceAll(string(server), "\n", ""),
			Username: strings.ReplaceAll(string(username), "\n", ""),
			Secret:   strings.ReplaceAll(string(secret), "\n", ""),
		},
	}

	data.Create()
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

	config, err := os.Open(fmt.Sprintf("%s/gitconfig", util.GetHome()))
	if err != nil {
		// TODO: create new one
		log.Fatalln(err)
		return
	}
	defer config.Close()

	credential, err := os.Open(fmt.Sprintf("%s/.git-credentials", util.GetHome()))
	if err != nil {
		// TODO: create new one
		log.Fatalln(err)
		return
	}
	defer credential.Close()

	// TODO: create change config

	uri := fmt.Sprintf("%s://%s:%s@%s",
		prof.AuthData.Protocol,
		prof.AuthData.Username,
		prof.AuthData.Secret,
		prof.AuthData.Server)
	_, err = credential.Write([]byte(uri))
	if err != nil {
		return
	}
}
