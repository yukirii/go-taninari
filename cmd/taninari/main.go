package main

import (
	"fmt"
	"os"

	"github.com/shiftky/go-taninari"
	"github.com/urfave/cli"
)

const cliVersion = "0.1.0"

func Show(goroku *taninari.Goroku) {
	fmt.Println("たになり語録 - " + goroku.PublishedAt)
	fmt.Println(goroku.Text)
	fmt.Println(goroku.PublishedURL)
}

func main() {
	app := cli.NewApp()
	app.Name = "taninari"
	app.Usage = "人生楽しんでますか？"
	app.Version = cliVersion
	app.Action = func(c *cli.Context) error {
		goroku, _ := taninari.GetRandomGoroku()
		Show(goroku)
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "patriot",
			Usage: "launch a missile",
			Action: func(c *cli.Context) error {
				ShowPatriot()
				return nil
			},
		},
	}
	app.Run(os.Args)
}
