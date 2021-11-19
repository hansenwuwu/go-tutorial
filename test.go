package main

import (
	"flag"
	"fmt"

	"github.com/hansenwuwu/go-tutorial/myMod/pk1"
)

var (
	isSandBox = flag.Bool("isSandBox", false, "If the env is Prod")
)

func main() {
	flag.Parse()
	fmt.Println(*isSandBox)
	pk1.SayMyMod()
}
