package main

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"bufio"
	"strings"
	"sort"
)
var sc = bufio.NewScanner(os.Stdin)

func List(c *cli.Context){
	fmt.Println("メンターリスト")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名前", "学部", "学科","優先順位","業界"})
	for _, v := range GetMentors() {
		str := ""
		for i,in := range v.Industries{
			str += strconv.Itoa(i)+":"+in.Name+" "
		}
		table.Append([]string{v.Name,v.Faculty,v.Department,strconv.Itoa(v.Priority),str})
	}
	table.Render()
}

func Match(c *cli.Context){
	ent := insertEnter()
	mentor := GetMentors()
	var matchMentors []MatchMentor
	for _,v := range mentor{
		matchMentors = append(matchMentors,calculateMentor(ent,v))
	}
	sort.Slice(matchMentors,func(i,j int)bool{
		return matchMentors[i].Score > matchMentors[j].Score
	})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名前", "SCORE"})
	for _, v := range matchMentors {
		table.Append([]string{v.Name,strconv.Itoa(v.Score)})
	}
	table.Render()
}

func separateIndustry(str string)[]Industry{
	in := strings.Split(str," ")
	var industries []Industry
	for _,v := range in {
		industries = append(industries,Industry{Name:v})
	}
	return industries
}

func insertEnter()Enter{
	var ent Enter
	fmt.Println("失敗した場合はcommand + C")
	fmt.Print("エンター名を入力してください")
	if sc.Scan() {
		ent.Name = sc.Text()
	}
	fmt.Print("学部を入力してください")
	if sc.Scan() {
		ent.Faculty = sc.Text()
	}
	fmt.Print("学科を入力してください")
	if sc.Scan() {
		ent.Department = sc.Text()
	}
	fmt.Print("紹介者を入力してください(いなければ”なし”)")
	if sc.Scan() {
		ent.Introducer = sc.Text()
	}
	fmt.Println("志望業界を入力してください(複数の場合は半角スペースで)")
	if sc.Scan() {
		ent.Industries = separateIndustry(sc.Text())
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名前", "学部", "学科","紹介者","業界"})
	str := ""
	for i,in := range ent.Industries{
		str += strconv.Itoa(i)+":"+in.Name+" "
	}
	table.Append([]string{ent.Name,ent.Faculty,ent.Department,ent.Introducer,str})
	table.Render()
	fmt.Println("これでよろしいですか y/n")
	if sc.Scan() {
		if strings.Index(sc.Text(),"y") == -1{
			fmt.Println("やり直してください")
			os.Exit(0)
		}
	}
	return ent
}

func calculateMentor(ent Enter,ment Mentor)MatchMentor{
	score := 0
	ratio := GetRatio()
	if strings.Index(ent.Faculty,ment.Faculty) != -1 {
		score += 1*ratio.Faculty
	}
	if strings.Index(ent.Department,ment.Department) != -1 {
		score += 1*ratio.Department
	}
	if strings.Index(ent.Introducer,ment.Name) != -1 {
		score += 1*ratio.Faculty
	}
	for _,v := range ent.Industries{
		score += contains(ment.Industries,v.Name) * ratio.Industries
	}
	return MatchMentor{Name:ment.Name,Score:score}
}

func contains(in []Industry, str string) int {
	for _, v := range in {
		if strings.Index(v.Name, str) != -1 {
			return 1
		}
	}
	return 0
}