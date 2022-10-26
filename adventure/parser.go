package adventure

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Saves the dynamic keys (chapter names) as keys for the map.
type Adventure map[string]AdvPart

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

func (a Adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var chapter string
	if r.URL.Path != "/favicon.ico" && r.URL.Path != "/css/myCss.css" {
		fmt.Println(r.URL.Path)
		chapter = strings.Trim(r.URL.Path, "/")
	}
	t, err := template.ParseFiles("../html/templates/adventure.html")
	if err != nil {
		log.Fatal("Can't parse the files. Error: ", err)
	}

	if a[chapter].Title == "" {
		fmt.Fprint(w, "Please direct to 'ler0y.com:8081/intro' to start your adventure")
	} else {
		err = t.Execute(w, a[chapter])
		if err != nil {
			log.Fatal("Can't execute the templates. Error: ", err)
		}
	}
}

// Unmarshals the data for the specified chapter into adv and returns it.
func JSONParser(path string, chapter string) Adventure {

	var adv Adventure

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
	return adv
}
