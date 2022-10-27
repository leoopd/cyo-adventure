package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leoopd/cyo-adventure/adventure"
)

func main() {

	file := "gopher.json"
	chapter := "intro"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "To use the adventure, please direct to ler0y.com:8081/intro")
	})

	adv := adventure.JSONParser(file, chapter)
	
	http.Handle("/style.css", http.FileServer(http.Dir("../html/templates/style.css")))
	http.Handle("/intro", adv)

	log.Print("Listening on :8081")
	err := http.ListenAndServe(":8081", adv)
	if err != nil {
		log.Fatal(err)
	}

}
