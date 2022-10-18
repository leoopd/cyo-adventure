package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type adventure struct {
	Title   string    `json:"title"`
	Options []Options `json:"options"`
	Story   []string  `json:"story"`
}

type Options struct {
	Text map[string]string `json:"text"`
	Arc  map[string]string `json:"arc"`
}

type advOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONParser(jsonFile string) {

	defer os.Exit(1)

	// var adventure adventure
	var adventureMap map[string]adventure

	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully read", jsonFile)
	json.Unmarshal(file, &adventureMap)
	fmt.Printf("%#+v\n", adventureMap)
}
