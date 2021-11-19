package main

import (
	"flag"
	"fmt"
)

var (
	isSandBox = flag.Bool("isSandBox", false, "If the env is Prod")
)

func main() {
	flag.Parse()
	fmt.Println(*isSandBox)
}
