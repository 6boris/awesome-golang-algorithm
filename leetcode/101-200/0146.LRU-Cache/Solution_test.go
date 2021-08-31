package Solution

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(2)
	fmt.Println(lru.Get(2))

	lru.Put(2, 6)
	fmt.Println(lru.Get(1))

	lru.Put(1, 5)
	lru.Put(1, 2)

	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))

}
