package main

import (
	"fmt"
	"os"

	"github.com/shiftky/go-taninari"
	"github.com/urfave/cli"
)

const cliVersion = "0.2.0"

func Show(message *taninari.GorokuMessage) {
	fmt.Println("たになり語録 - " + message.PublishedAt)
	fmt.Println(message.Text)
	fmt.Println(message.PublishedURL)
}

func main() {
	app := cli.NewApp()
	app.Name = "taninari"
	app.Usage = "人生楽しんでますか？"
	app.Version = cliVersion
	app.Action = func(c *cli.Context) error {
		message, _ := taninari.GetRandomMessage()
		Show(message)
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

				messages, _ := taninari.SearchMessages(c.Args().Get(0))

				cnt := len(messages)
				if cnt > 0 {
					fmt.Println(cnt, "個のメッセージがみつかりましたね。")
					for i, message := range messages {
						fmt.Println("\n\x1b[32m", i+1, message.Text, "\x1b[0m")
						fmt.Println(message.PublishedURL, "-", message.PublishedAt)
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
