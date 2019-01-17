package main

import (
	"fmt"
	"testing"
)

func TestRequest(t *testing.T) {
	resp := Request("GET", GITHUB_CONTRIBUTOR_API_URL, nil)
	fmt.Println(string(resp))
}
