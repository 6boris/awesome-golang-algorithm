package leetcode

import (
	"testing"
)

func TestUrlPath(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		err := UrlPath("https://www.baidu.com")
		if err != nil {
			t.Fatal(err.Error())
		}
	})
}
