package main

type MentorConfig struct {
	Mentors []Mentor
}

type Mentor struct {
	Name string
	Faculty string
	Department string
	Priority  int
	Industries []Industry
}

type Enter struct {
	Name string
	Faculty string
	Department string
	Industries []Industry
	Introducer string
}

type MatchMentor struct {
	Name string
	Score int
}

type Industry struct {
	Name string
}

type Ratio struct {
	Faculty int
	Department int
	Industries int
	Priority int
	Introducer int
}