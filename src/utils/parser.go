package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type contents struct {
	Part advPart `json:"intro"`
}

type advPart struct {
	Title   string `json:"title"`
	Options []advOptions
	Story   []string `json:"story"`
}

type advOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONParser(jsonFile string) {

	defer os.Exit(1)

	var adv contents

	file, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully read", jsonFile)
	err = json.Unmarshal(file, &adv)
	fmt.Println(string(file))

	fmt.Println(err)
	fmt.Printf("%#+v\n", adv.Part.Story)
}
