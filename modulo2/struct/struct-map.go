// Go in action
// @jeffotoni
// 2019-05-06
// different ways to inizialize a struct

package main

import "fmt"

type Before struct {
	m map[string]string
}

func contrivedAfter(b interface{}) interface{} {
	return struct {
		Before
		s []string
	}{b.(Before), []string{"new value", "value two"}}
}

func main() {
	b := Before{map[string]string{"some": "value"}}
	a := contrivedAfter(b)
	fmt.Println(a)
}
