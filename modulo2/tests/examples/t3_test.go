// Go in action
// @jeffotoni
// 2019-05-06

package main

import "testing"

func TestSum(t *testing.T) {
	x := 1 + 1
	if x != 11 {
		t.Error("Error! 1 + 1 is not equal to 2, I got", x)
	}
}
