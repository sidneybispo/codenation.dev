// Go in action
// @jeffotoni
// 2019-05-06
package main

import "fmt"

import "pkg/math"
import log "github.com/sirupsen/logrus"
import . "github.com/jeffotoni/gcolor"

func main() {
	fmt.Println("hello teste workspaces!!!")

	log.WithFields(log.Fields{
		"curso":  "codenation",
		"number": 1,
		"size":   100,
	}).Info("A walrus appears")

	Yellow.Cprintln("test pacote gcolor GOPATH!")

	fmt.Println("Math:", math.Sum(3, 7))
}
