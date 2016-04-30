package main

import (
	"github.com/codegangsta/cli"
)

var gName string
var gTemplate string
var gCategory string

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
			Destination: &gName,
		},
		cli.StringFlag{
			Name:        "template, t",
			Usage:       "using text template",
			Destination: &gTemplate,
		},
		cli.BoolFlag{
			Name:   "open, o",
			Usage:  "open text editor",
			EnvVar: "DEFAULT",
		},
		cli.StringFlag{
			Name:        "category, c",
			Usage:       "make memo category",
			Destination: &gCategory,
		}},
}
