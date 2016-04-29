package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/codegangsta/cli"
)

func newInit(c *cli.Context) {
	if c.NArg() <= 0 {
		fmt.Println("no firectory name")
		return
	}
	dirName := c.Args()[0]
	if err := os.MkdirAll(dirName+"/memo", 0777); err != nil {
		fmt.Println(err)
	}
	if err := os.MkdirAll(dirName+"/template", 0777); err != nil {
		fmt.Println(err)
	}
	config := newConfig()
	config.Name = dirName
	config.marshalJSON()

	fmt.Println("create " + dirName + " directory")
}

func newAction(c *cli.Context) {
	if false == isExist("memo") {
		fmt.Println("failed! memobase directory is not exist")
		return
	}

	config := unmarshalJSON()
	filename := createFilename(time.Now().Format(config.NewFileDatetime))

	if len(name) > 0 {
		filename = createFilename(name)
	}

	if len(category) > 0 {
		if err := os.MkdirAll("memo/"+category, 0777); err != nil {
			fmt.Println("category create failed")
		}
		filename = category + "/" + filename
	}
	dst, err := os.Create("memo/" + filename)
	if err != nil {
		fmt.Println("file create failed")
		os.Exit(1)
	}

	if len(template) > 0 {
		templatePath := "template/" + template
		if false == isExist(templatePath) {
			fmt.Println("failed! template directory is not exist")
		} else {
			src, err := os.Open(templatePath)
			if err != nil {
				panic(err)
			}
			defer src.Close()
			_, err = io.Copy(dst, src)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("create %s\n", filename)
	if c.Bool("open") == true {
		cmd := exec.Command(config.Editor, "memo/"+filename)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
