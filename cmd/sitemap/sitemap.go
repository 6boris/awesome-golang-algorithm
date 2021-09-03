package sitemap

import (
	"fmt"

	"github.com/kylesliu/awesome-golang-algorithm/cmd/leetcode"
)

type Config struct {
	SecureKey string
}

type SiteMap struct {
	UrlSet []UrlSet `xml:"urlset"`
}

func New(problems []leetcode.Problem) SiteMap {
	fmt.Println(len(problems))
	for i := 0; i < len(problems); i++ {
		fmt.Println(problems[i].DirPath)
	}
	return SiteMap{}
}

type UrlSet struct {
	Url []Url `xml:"url"`
}

type Url struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}
