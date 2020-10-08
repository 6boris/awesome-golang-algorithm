package Solution

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	emailSet := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		localName := strings.Replace(strings.Split(parts[0], "+")[0], ".", "", -1)
		emailSet[localName+"@"+parts[1]] = true
	}
	return len(emailSet)
}
