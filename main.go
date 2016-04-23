package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "memobase"
	app.Usage = "memo manage app by make golang"
	app.Version = "0.0.2"
	app.Commands = Commands
	app.Run(os.Args)
}
