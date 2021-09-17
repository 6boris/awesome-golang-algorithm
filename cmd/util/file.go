package util

import (
	"fmt"
	"os"
)

func WriteFile(path, s string) {
	f, err := os.Create("sitemap.xml")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	f.WriteString(s)
}
