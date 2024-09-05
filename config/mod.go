package config

import (
	"encoding/json"
	"fmt"
	"github.com/devproje/git-aman/util"
	"github.com/devproje/plog/log"
	"os"
)

type Config struct {
	SetId  int `json:"set_id"`
	LastId int `json:"last_id"`
}

var Conf *Config

func Load() *Config {
	file := getConfigFile()
	f, _ := os.ReadFile(file)
	var raw *Config

	err := json.Unmarshal(f, &raw)
	if err != nil {
		raw = &Config{
			SetId:  0,
			LastId: 1,
		}
	}

	return raw
}

func SetId(id int) {
	Conf.SetId = id
	save()
}

func SetLastId(id int) {
	Conf.LastId = id
	save()
}

func save() {
	file := getConfigFile()
	pak, _ := json.Marshal(Conf)
	err := os.WriteFile(file, pak, 0644)
	if err != nil {
		return
	}
}

func getConfigFile() string {
	dir := util.GetDataDir()
	file := fmt.Sprintf("%s/config.json", dir)
	if _, err := os.Stat(file); err != nil {
		if !os.IsNotExist(err) {
			log.Panicln(err)
		}

		_, err = os.Create(file)
		if err != nil {
			log.Panicln(err)
		}
	}

	return file
}
