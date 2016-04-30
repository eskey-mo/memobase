package main

import (
	"fmt"
	"os"
	"os/exec"
)

func createFilename(filename string) string {
	return filename + ".txt"
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func makeDirectory(dirname string) {
	if err := os.MkdirAll(dirname, 0777); err != nil {
		fmt.Println(dirname + " create failed")
		fmt.Println(err)
		os.Exit(1)
	}
}

func createFile(filename string) *os.File {
	dst, err := os.Create(filename)
	if err != nil {
		fmt.Println(filename + " create failed")
		os.Exit(1)
	}
	return dst
}

func execCommand(cmd string, args string) {
	exec := exec.Command(cmd, args)
	exec.Stdin = os.Stdin
	exec.Stdout = os.Stdout
	exec.Run()
}
