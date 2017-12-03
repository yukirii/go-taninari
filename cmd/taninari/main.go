package main

import (
	"fmt"
	"os"

	"github.com/shiftky/go-taninari"
	"github.com/urfave/cli"
)

const cliVersion = "0.2.0"

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
			Name:      "search",
			Usage:     "search Taninari's messages",
			ArgsUsage: "[search keyword]",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					cli.ShowSubcommandHelp(c)
					return nil
				}

				fmt.Println("Keyword: " + c.Args().Get(0))

				gorokus, _ := taninari.SearchGorokus(c.Args().Get(0))

				cnt := len(gorokus)
				if cnt > 0 {
					fmt.Println(cnt, "個のメッセージがみつかりましたね。")
					for i, goroku := range gorokus {
						fmt.Println("\n\x1b[32m", i+1, goroku.Text, "\x1b[0m")
						fmt.Println(goroku.PublishedURL, "-", goroku.PublishedAt)
					}
				} else {
					fmt.Println("むむむ。何も見つからなかったみたいですね。")
				}

				return nil
			},
		},
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
