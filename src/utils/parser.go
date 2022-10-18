package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type adventure struct {
	Arc string `json:"intro"`
}

type advBody struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []advOption `json:"options"`
}

type advOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONParser(jsonFile string) {

	defer os.Exit(1)

	// var adventure adventure
	var adventureMap map[string]interface{}

	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully read", jsonFile)
	json.Unmarshal(file, &adventureMap)
	fmt.Printf("%#+v\n", adventureMap)
}
