//+build !test

package main

import (
	"fmt"
	"os"
)

func main() {
	err := outCommand(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
