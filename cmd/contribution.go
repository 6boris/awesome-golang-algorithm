package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
)

type Contributor struct {
	Login            string   `json:"login"`
	Avatar_Url       string   `json:"avatar_url"`
	Name             string   `json:"name"`
	Url              string   `json:"url"`
	Profile          string   `json:"profile"`
	Contributions    int      `json:"contributions"`
	ContributionType []string `json:"contribution_type"`
}

const GITHUB_CONTRIBUTOR_API_URL = "https://api.github.com/repos/kylesliu/awesome-golang-algorithm/contributors"
const GITHUB_CONTRIBUTOR_TMPL_PATH = "./tpl/.all-contributorsrc"

func getContributorBufer() []byte {

	contributor_buffer := Request("GET", GITHUB_CONTRIBUTOR_API_URL, nil)

	return contributor_buffer
}

func GetContributorString() string {
	str := string(getContributorBufer())
	return str
}

//	get the josnmarshal json
func GetContributorJosnMarshal(prefix, indent string) string {
	contributors := []Contributor{}

	if err := json.Unmarshal(getContributorBufer(), &contributors); err != nil {
		log.Fatalln(err.Error())
	}

	contributors_buffer, err := json.MarshalIndent(contributors, prefix, indent)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(contributors_buffer)
}

func GetContributorInstance() []Contributor {
	contributors := []Contributor{}

	if err := json.Unmarshal(getContributorBufer(), &contributors); err != nil {
		log.Fatalln(err.Error())
	}
	return contributors
}
func getContributorTemplate() string {
	buffer := ReadFile(GITHUB_CONTRIBUTOR_TMPL_PATH)
	return string(buffer)
}

func GenerateContributorTemplete() {
	tpl_str := getContributorTemplate()
	//fmt.Println(tpl_str)
	contributors := GetContributorInstance()

	fmt.Println(contributors)
	//	Tmpl
	var tmpRes bytes.Buffer

	tmpl, err := template.New("Contributors: ").Parse(tpl_str)
	if err != nil {
		log.Printf("%s", err)
	}
	err = tmpl.Execute(&tmpRes, contributors)
	fmt.Println(string(tmpRes.Bytes()))
}
