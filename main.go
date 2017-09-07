package main

import (
	"fmt"
	"os"

	"gitlab.com/Lesterpig/invoice/manager"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "new" {
		filename, err := manager.Next()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(filename)
		os.Exit(0)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "[input yaml]", "[output pdf]")
		fmt.Fprintln(os.Stderr, "      ", os.Args[0], "new")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	err := manager.Generate(inputFile, outputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
