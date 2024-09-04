package profile

import (
	"encoding/json"
	"fmt"
	"github.com/devproje/git-aman/types"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
)

type Profile struct {
	Id          int        `json:"id"`
	DisplayName string     `json:"display_name"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	AuthData    Credential `json:"credential"`
}

type Credential struct {
	Protocol types.Protocol `json:"protocol"`
	Server   string         `json:"server"`
	Username string         `json:"username"`
	Secret   string         `json:"secret"`
}

var profs []Profile

func (prof *Profile) Create() {
	profs = append(profs, *prof)

	save()
}

func Read(id int) *Profile {
	for _, prof := range profs {
		if prof.Id == id {
			return &prof
		}
	}

	return nil
}

func Delete(id int) {
	for i, prof := range profs {
		if prof.Id == id {
			profs = append(profs[:i], profs[i+1:]...)
			break
		}
	}

	save()
}

func QueryProfs() {
	for _, prof := range profs {
		fmt.Printf("profile_id: %d\n", prof.Id)
		fmt.Printf("\tdisplay_name: %s\n", prof.DisplayName)
		fmt.Printf("\tuser: %s\n", prof.Name)
		fmt.Printf("\temail: %s\n", prof.Email)
	}

	log.Printf("total profile size: %d\n", len(profs))
}

func LoadAll() {
	dir := loadPath()
	f, _ := os.ReadFile(fmt.Sprintf("%s/profile.json", dir))
	var data []Profile

	err := json.Unmarshal(f, &data)
	if err != nil {
		profs = make([]Profile, 0)
		return
	}

	profs = data
}

func ProfSize() int {
	return len(profs)
}

func save() {
	dir := loadPath()
	data, _ := json.Marshal(profs)
	err := os.WriteFile(fmt.Sprintf("%s/profile.json", dir), data, 0644)
	if err != nil {
		log.Errorln(err)
	}
}

func loadPath() string {
	dir := util.GetDataDir()
	file := fmt.Sprintf("%s/profile.json", dir)
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
