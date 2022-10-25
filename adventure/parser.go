package adventure

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Saves the dynamic keys (chapter names) as keys for the map.
type chapters map[string]AdvPart

// Structs to store the chapter data that gets Unmarshaled.
type AdvPart struct {
	Title   string       `json:"title"`
	Options []AdvOptions `json:"options"`
	Story   []string     `json:"story"`
}

type AdvOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (a AdvPart) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	t, err := template.ParseFiles("./templates/adventure.html")
	if err != nil {
		log.Fatal("Can't parse the files. Error: ", err)
	}
	err = t.Execute(w, a)
	if err != nil {
		log.Fatal("Can't execute the templates. Error: ", err)
	}
}

// Unmarshals the data for the specified chapter into adv and returns it.
func JSONParser(path string, chapter string) AdvPart {

	var adv chapters

	// Reads a .json from the specified path.
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshals the file to adv
	err = json.Unmarshal(file, &adv)
	if err != nil {
		log.Fatal(err)
	}
	return adv[chapter]
}
