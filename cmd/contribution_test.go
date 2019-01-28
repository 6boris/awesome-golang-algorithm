package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetContributorString(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}

	str := GetContributorString()
	fmt.Println(str)
}

func TestGetContributorJosnMarshal(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}

	str := GetContributorJosnMarshal("", "	")
	fmt.Println(str)
}

func TestGetContributorInstance(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}

	contributors := GetContributorInstance()
	fmt.Println(contributors)
}

func TestGenerateContributorTemplete(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}

	GenerateContributorTemplete()
}

func BenchmarkGetContributorInstance(b *testing.B) {
	if runtime.GOOS != "windows" {
		return
	}

	b.N = 2000000
	contributors := GetContributorInstance()
	fmt.Println(contributors)
}
