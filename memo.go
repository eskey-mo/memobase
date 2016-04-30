package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

const (
	// MEMO is memo directory name
	MEMO = "memo"
	// TEMPLATE is template directory name
	TEMPLATE = "template"
)

func newInit(c *cli.Context) {
	if c.NArg() <= 0 {
		fmt.Println("no firectory name")
		return
	}
	dirname := c.Args()[0]
	makeDirectory(dirname + "/" + MEMO)
	makeDirectory(dirname + "/" + TEMPLATE)
	config := newConfig()
	config.Name = dirname
	config.marshalJSON()

	fmt.Println("create " + dirname + " directory")
}

func newAction(c *cli.Context) {
	// memo exist
	if false == isExist(MEMO) {
		fmt.Println("failed! memobase directory is not exist")
		return
	}

	// read config file
	config := unmarshalJSON()
	filename := createFilename(time.Now().Format(config.NewFileDatetime))

	// name option
	if len(gName) > 0 {
		filename = createFilename(gName)
	}

	// category option
	if len(gCategory) > 0 {
		filename = optionCategoryByNewAction(filename, gCategory)
	}

	// create new memo file
	fileDirectory := createFile(MEMO + "/" + filename)

	// template option
	if len(gTemplate) > 0 {
		optionTemplateByNewAction(fileDirectory, gTemplate)
	}

	// open option
	if c.Bool("open") == true {
		execCommand(config.Editor, MEMO+"/"+filename)
	}
	// command finish
	fmt.Printf("create %s\n", filename)
}

func optionCategoryByNewAction(filename string, category string) string {
	makeDirectory(MEMO + "/" + category)
	return category + "/" + filename
}

func optionTemplateByNewAction(filedirectory *os.File, template string) {
	templatePath := TEMPLATE + "/" + template
	if false == isExist(templatePath) {
		fmt.Println("failed! template directory is not exist")
	} else {
		src, err := os.Open(templatePath)
		if err != nil {
			panic(err)
		}
		defer src.Close()
		if _, err = io.Copy(filedirectory, src); err != nil {
			panic(err)
		}
	}
}
