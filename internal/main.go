package main

import (
	"flag"
	"fmt"
	"os"
	"recipes/utilities/initialization"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	// This is where the Code Will Start
	flag.Usage = usage
	settings := flag.String("S", "", "Set your settings")
	flag.Parse()
	
	// Execute Initials
	initialization.InitializationExecute(*settings)

}
