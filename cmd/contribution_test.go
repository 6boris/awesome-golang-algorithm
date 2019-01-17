package main

import (
	"fmt"
	"testing"
)

func TestGetContributorString(t *testing.T) {
	str := GetContributorString()
	fmt.Println(str)
}

func TestGetContributorJosnMarshal(t *testing.T) {
	str := GetContributorJosnMarshal("", "	")
	fmt.Println(str)
}

func TestGetContributorInstance(t *testing.T) {
	contributors := GetContributorInstance()
	fmt.Println(contributors)
}

func TestGenerateContributorTemplete(t *testing.T) {
	GenerateContributorTemplete()
}

func BenchmarkGetContributorInstance(b *testing.B) {
	b.N = 2000000
	contributors := GetContributorInstance()
	fmt.Println(contributors)
}
