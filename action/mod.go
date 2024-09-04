package action

import (
	"bufio"
	"fmt"
	"github.com/devproje/git-aman/profile"
	"github.com/devproje/git-aman/types"
	"golang.org/x/term"
	"strings"
	"syscall"
)

func input(msg string, reader *bufio.Reader) string {
	fmt.Printf(msg)
	buf, _ := reader.ReadBytes('\n')

	return strings.ReplaceAll(string(buf), "\n", "")
}

func setDisplayName(prof *profile.Profile, reader *bufio.Reader) {
	prof.DisplayName = input("type your profile display name: ", reader)
}

func setConfigName(prof *profile.Profile, reader *bufio.Reader) {
	prof.Name = input("type git config 'user.name': ", reader)
}

func setConfigEmail(prof *profile.Profile, reader *bufio.Reader) {
	prof.Email = input("type git config 'user.email': ", reader)
}

func setProtocol(prof *profile.Profile, reader *bufio.Reader) {
	raw := types.CheckProto(input("type your git protocol: ", reader))
	prof.AuthData.Protocol = types.CheckProto(string(raw))
}

func setServer(prof *profile.Profile, reader *bufio.Reader) {
	prof.AuthData.Server = input("type your git server: ", reader)
}

func setUsername(prof *profile.Profile, reader *bufio.Reader) {
	prof.AuthData.Username = input("type your git username: ", reader)
}

func setSecret(prof *profile.Profile) {
	fmt.Printf("type git secret: ")
	secret, _ := term.ReadPassword(syscall.Stdin)

	prof.AuthData.Secret = strings.ReplaceAll(string(secret), "\n", "")
}

func parser(cmd string, prof *profile.Profile, reader *bufio.Reader) {
	switch cmd {
	case "display":
		setDisplayName(prof, reader)
		break
	case "name":
		setConfigName(prof, reader)
		break
	case "email":
		setConfigEmail(prof, reader)
		break
	case "protocol":
		setProtocol(prof, reader)
		break
	case "server":
		setServer(prof, reader)
		break
	case "username":
		setUsername(prof, reader)
		break
	case "secret":
		setSecret(prof)
		break
	case "status":
		status(prof)
	case "quit":
		syscall.Exit(0)
	case "?":
		fmt.Printf("[available command]\n")
		fmt.Println("display, name, email, protocol, server, username, secret, status, quit")
	default:
		fmt.Printf("unknown command: %s\n", cmd)
	}
}

func status(prof *profile.Profile) {
	fmt.Printf("[General]\n")
	fmt.Printf("\tdisplay_name: %s\n", prof.DisplayName)
	fmt.Printf("[GitConfig]\n")
	fmt.Printf("\tname: %s\n", prof.Name)
	fmt.Printf("\temail: %s\n", prof.Email)
	fmt.Printf("[AuthData]\n")
	fmt.Printf("\tproto: %s\n", prof.AuthData.Protocol)
	fmt.Printf("\tserver: %s\n", prof.AuthData.Server)
	fmt.Printf("\tusername: %s\n", prof.AuthData.Username)
	fmt.Printf("\tsecret: [hidden]\n")
	fmt.Printf("if you're typed info is correct, please type 'y' in console\n")
}
