package main

import (
	"log"
	"net/http"

	"github.com/leoopd/cyo-adventure/adventure"
)

type varss struct {
	Greeting string
	Name     string
}

func main() {

	file := "gopher.json"
	chapter := "intro"

	adv := adventure.JSONParser(file, chapter)
	http.Handle("/intro", adv)

	log.Print("Listening on :8081")
	err := http.ListenAndServe(":8081", adv)
	if err != nil {
		log.Fatal(err)
	}
}
