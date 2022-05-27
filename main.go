package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed help.txt
var helpText string

func main() {
	var helpFlag bool;
	
	flag.BoolVar(&helpFlag, "help", false, "Shows the help page")
	flag.Parse()

	if (helpFlag || len(os.Args) > 1) {
		fmt.Println(helpText)
		return
	}
}