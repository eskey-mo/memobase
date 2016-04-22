package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

var name string
var template string

func main() {
	app := cli.NewApp()
	app.Name = "memobase"
	app.Usage = "memo manage app for golang"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new memo",
			Action:  newAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name, n",
					Usage:       "create filename",
					Destination: &name,
				},
				cli.StringFlag{
					Name:        "template, t",
					Usage:       "using text template",
					Destination: &template,
				}},
		},
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}

func newAction(c *cli.Context) {
	filename := time.Now().Format("2006-01-02-15-04-05") + ".txt"

	if len(name) > 0 {
		filename = name + ".txt"
	}

	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println("ファイルが作成できませんでした。")
		os.Exit(1)
	}

	if len(template) > 0 {
		src, err := os.Open(template)
		if err != nil {
			panic(err)
		}
		defer src.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("create %s\n", filename)
}
