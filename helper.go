package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var path string

func init() {
	p, _ := os.Executable()
	path = filepath.Dir(p)
}

func GetMentors() []Mentor {
	var config MentorConfig
	_, err := toml.DecodeFile(path+"/mentor.toml", &config)
	if err != nil {
		panic(err)
	}
	return config.Mentors
}

func GetEnter() Enter {
	var config Enter
	_, err := toml.DecodeFile(path+"/enter.toml", &config)
	if err != nil {
		panic(err)
	}
	return config
}

func GetRatio() Ratio {
	var config Ratio
	_, err := toml.DecodeFile(path+"/ratio.toml", &config)
	if err != nil {
		panic(err)
	}
	return config
}
