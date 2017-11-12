//+build !test

package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()
	err := inCommand(fs, os.Args, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
