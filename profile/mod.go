package profile

import (
	"encoding/json"
	"fmt"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
	"path/filepath"
)

type Profile struct {
	id          int
	DisplayName string
	Name        string
	Email       string
	AuthData    Credential
}

type Credential struct {
	Server   string
	Username string
	Secret   string
}

var profs []Profile

func (prof *Profile) Create(id int) {
	prof.id = id
	profs = append(profs, *prof)

	save()
}

func Read(id int) *Profile {
	for _, prof := range profs {
		if prof.id == id {
			return &prof
		}
	}

	return nil
}

func Delete(id int) {
	for i, prof := range profs {
		if prof.id == id {
			profs = append(profs[:i], profs[i+1:]...)
			break
		}
	}

	save()
}

func QueryProfs() {
	for _, prof := range profs {
		log.Println("id: %d", prof.id)
		fmt.Printf("\t%s\n", prof.DisplayName)
		fmt.Printf("\t%s\n", prof.Name)
		fmt.Printf("\t%s\n", prof.Email)
	}

	log.Printf("total profile size: %d\n", len(profs))
}

func LoadAll() {
	file := loadPath()
	f, _ := os.ReadFile(file)
	var data []Profile

	err := json.Unmarshal(f, &data)
	if err != nil {
		profs = make([]Profile, 0)
		return
	}

	profs = data
}

func save() {
	file := loadPath()
	data, _ := json.Marshal(profs)
	err := os.WriteFile(file, data, 0644)
	if err != nil {
		log.Errorln(err)
	}
}

func loadPath() string {
	dir := util.GetDataDir()
	file := filepath.Join(dir, "profile.json")
	if _, err := os.Stat(file); err != nil {
		if !os.IsNotExist(err) {
			log.Panicln(err)
		}

		_, err = os.Create(file)
		if err != nil {
			log.Panicln(err)
		}
	}

	return dir
}
