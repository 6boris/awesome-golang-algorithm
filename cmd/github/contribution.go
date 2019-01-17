package github

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Contributor struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
	Url        string `json:"url"`
}

const GITHUB_CONTRIBUTOR_API_URL = "https://api.github.com/repos/kylesliu/awesome-golang-leetcode/contributors"
const GITHUB_CONTRIBUTOR_TMPL_PATH = "cmd/template/contributors/.all-contributorsrc"

func getContributorBufer() []byte {
	request, err := http.Get(GITHUB_CONTRIBUTOR_API_URL)

	if err != nil {
		log.Panicln("Lettcode Problem 接口获取失败：", err)
	}

	if request.StatusCode != 200 {
		log.Panicln("Lettcode Problem 接口地址不存在：", err)
	}

	body, _ := ioutil.ReadAll(request.Body)
	return body
}

func GetContributorString() string {
	str := string(getContributorBufer())
	return str
}
