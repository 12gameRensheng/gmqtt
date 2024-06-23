package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type UserConfigType struct {
	Users []User `yaml:"users"`
}

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var UserConfig = UserConfigType{}

func (userConfig *UserConfigType) ParseConfig(filePath string) {
	if filePath == "" {
		return
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &userConfig)
	if err != nil {
		return
	}
}
