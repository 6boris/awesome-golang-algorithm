package Solution

import (
	"strings"
)

func Solution(text string) string {
	entities := []string{"&quot;", "&apos;", "&gt;", "&lt;", "&frasl;", "&amp;"}
	entityMap := []string{"\"", "'", ">", "<", "/", "&"}
	for index, e := range entities {
		text = strings.ReplaceAll(text, e, entityMap[index])
	}
	return text
}
