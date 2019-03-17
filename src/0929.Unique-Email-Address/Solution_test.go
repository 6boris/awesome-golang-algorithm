package Solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	listTest := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	got := numUniqueEmails(listTest)
	want := 2
	if got != want {
		t.Errorf("want '%d' but got '%d'", want, got)
	}
}
