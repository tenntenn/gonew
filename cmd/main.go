package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[0], flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd string, args []string) error {
	return nil
}
