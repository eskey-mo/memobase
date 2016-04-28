package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// config is application config
type config struct {
	Name   string `json:"name"`
	Editor string `json:"editor"`
}

// newConfig is config.json parse
func newConfig() *config {
	return &config{Name: "default", Editor: "vim"}
}

// marshalJSON is config.json parse to json
func (c *config) marshalJSON() {
	output, err := json.MarshalIndent(&c, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile(c.Name+"/config.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

// unmarshalJSON is config.json parse to struct
func unmarshalJSON() *config {
	// ファイルの読み込み
	rf, err := ioutil.ReadFile("./config.json") //[]byte型での読み込み
	// r, err := os.Open("./sample.json") //File型での読み込み
	if err != nil {
		log.Fatal(err)
	}
	data := new(config)
	if err := json.Unmarshal(rf, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}
	return data
}
