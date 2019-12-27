package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Octo Routin"
	app.Usage = "An example CLI for word counting"
	app.Author = "Chamod Perera"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "octo",
			Aliases: []string{"oc"},
			Usage:   "To count words count",
			Action: func() {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter text : ")
				text, _ := reader.ReadString('\n')
				fmt.Println(text)
			},
		},
	}
}

func main() {
	info()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
