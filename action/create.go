package action

import (
	"bufio"
	"github.com/devproje/git-aman/profile"
	"os"
)

func ProfileCreate() {
	var first = false
	var reader = bufio.NewReader(os.Stdin)
	var prof = profile.Profile{Id: profile.ProfSize() + 1}
	setDisplayName(&prof, reader)
	setConfigName(&prof, reader)
	setConfigEmail(&prof, reader)
	setProtocol(&prof, reader)
	setServer(&prof, reader)
	setUsername(&prof, reader)
	setSecret(&prof)

	for {
		if !first {
			status(&prof)
			first = true
		}

		cmd := input("[type y or ? ~]$ ", reader)
		if cmd == "y" {
			break
		}

		parser(cmd, &prof, reader)
	}

	prof.Create()
}
