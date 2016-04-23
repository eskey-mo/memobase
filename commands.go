package main

import (
	"github.com/codegangsta/cli"
)

var name string
var template string

// Commands list
var Commands = []cli.Command{
	commandInit,
	commandNew,
}
var commandInit = cli.Command{
	Name:   "init",
	Usage:  "create new memobase directory",
	Action: newInit,
}
var commandNew = cli.Command{
	Name:   "new",
	Usage:  "create new memo",
	Action: newAction,
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
}
