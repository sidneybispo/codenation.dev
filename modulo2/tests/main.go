// Go in action
// @jeffotoni
// 2019-05-06
//go run -ldflags "-X main.x=2 -X main.y=3" main.go

package main

import "strconv"

import (
	"github.com/jeffotoni/codenation.dev/aula01/examples/tests/pkg/math"
)

var (
	x, y   string
	xi, yi int
)

func init() {
	xi, _ = strconv.Atoi(x)
	yi, _ = strconv.Atoi(y)
}

func main() {
	println(math.Sum(xi, yi))
}
