package main

import (
	"fmt"

	"github.com/goodtoseeu57/easy-responses/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Print("initial work")
	}
}
