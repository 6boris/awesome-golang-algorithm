package Solution

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	emailSet := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		localName := strings.ReplaceAll(strings.Split(parts[0], "+")[0], ".", "")
		emailSet[localName+"@"+parts[1]] = true
	}
	return len(emailSet)
}
