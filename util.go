package main

import (
	"os"
)

func createFilename(filename string) string {
	return filename + ".txt"
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
