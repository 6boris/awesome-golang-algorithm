package Solution

import (
	"fmt"
	"testing"
)

//
// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	// ser := Constructor();
	deser := Constructor()
	// data := ser.serialize(root);
	ans := deser.deserialize("1,nil,nil")
	fmt.Println(ans)
}
