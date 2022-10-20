package main

import (
	"log"
	"net/http"

	"github.com/leoopd/cyo-adventure/adventure"
)

// type myHandler struct {
// 	Adventure utils.A
// }

// func (m myHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {

// }

func main() {

	file := "gopher.json"
	chapter := "intro"

	// for chapter != "home" {
	// 	chapter = adventure.Chapter(adventure.JSONParser(file, chapter))
	// }

	adv := adventure.JSONParser(file, chapter)
	http.Handle("/", adv)

	log.Print("Listening on :8081")
	err := http.ListenAndServe(":8081", adv)
	if err != nil {
		log.Fatal(err)
	}

}
