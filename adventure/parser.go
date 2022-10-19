package adventure

import (
	"encoding/json"
	"log"
	"os"
)

type Part struct {
	Intro     AdvPart `json:"intro"`
	NewYork   AdvPart `json:"new-york"`
	Debate    AdvPart `json:"debate"`
	SeanKelly AdvPart `json:"sean-kelly"`
	MarkBates AdvPart `json:"mark-bates"`
	Denver    AdvPart `json:"denver"`
	Home      AdvPart `json:"home"`
}

type AdvPart struct {
	Title   string       `json:"title"`
	Options []AdvOptions `json:"options"`
	Story   []string     `json:"story"`
}

type AdvOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JSONParser(path string) []byte {
	var adv Part

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &adv)
	return file
}
