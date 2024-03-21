package main

import (
	"fmt"

	"github.com/goodtoseeu57/easy-responses/cli"
)

func main() {
	fmt.Print("initial work")
	if err := cli.Execute(); err != nil {
		fmt.Print("it works")
	}
}
