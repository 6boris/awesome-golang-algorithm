package Solution

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	t.Run("test tree seilize", func(t *testing.T) {
		tree := conFromPreStr("124##5##36##")

		fmt.Println(dumpTreeToString(tree))
	})
}
