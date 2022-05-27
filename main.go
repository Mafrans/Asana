package main

import (
	_ "embed"
	"flag"
	"fmt"
	"mafrans/asana/api"
	"os"

	"github.com/joho/godotenv"
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

	loadDotEnv()

	meResponse := api.Routes.Me()
	fmt.Printf("Signed in as %s\n", meResponse.Data.Name)
}

func loadDotEnv() {
	err := godotenv.Load()
	if (err != nil) {
		panic("Error loading .env file")
	}
}