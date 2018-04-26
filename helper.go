package main

import "github.com/BurntSushi/toml"

func GetMentors()[]Mentor{
	var config MentorConfig
	_, err := toml.DecodeFile("mentor.toml", &config)
	if err != nil {
		panic(err)
	}
	return config.Mentors
}

func GetRatio()Ratio{
	var config Ratio
	_, err := toml.DecodeFile("ratio.toml", &config)
	if err != nil {
		panic(err)
	}
	return config
}