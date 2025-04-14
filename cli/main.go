package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

func main() {
	if len(os.Args) < 2 {
		color.Red.Println("Expecting a command. Run with help subcommand to view help.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--help", "help", "h":
		fmt.Printf("Help")

	case "phtml":
		if len(os.Args) != 3 {
			color.Red.Println("The phtml command expects a f8p file")
			os.Exit(1)
		}

		out, err := phtml(os.Args[2])
		if err != nil {
			color.Red.Println(err)
			os.Exit(1)
		}

		fmt.Println(out)
	}
}
