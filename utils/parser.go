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
	Title   string       `json:"title"`
	Options []advOptions `json:"options"`
	Story   []string     `json:"story"`
}

type advOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONParser(path string) (contents, []byte) {

	var adv contents

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &adv)

	return adv, file
}

func StoryTeller(c contents, json []byte) {

	for i := 0; i < len(c.Part.Story); i++ {
		fmt.Println(c.Part.Story[i])
	}

	fmt.Println()

	for j := 0; j < len(c.Part.Options); j++ {
		fmt.Printf("Option %d: %s\n", j+1, c.Part.Options[j].Text)
	}

}
