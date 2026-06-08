package util

import (
	"fmt"
	"os"
)

func WriteFile(path, s string) {
	f, err := os.Create("sitemap.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() { _ = f.Close() }()

	if _, err := f.WriteString(s); err != nil {
		fmt.Println(err.Error())
	}
}
