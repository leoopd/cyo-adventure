package main

import "github.com/leoopd/cyo-adventure/adventure"

func main() {

	file := "gopher.json"
	chapter := "intro"

	for chapter != "home" {
		chapter = adventure.Chapter(adventure.JSONParser(file, chapter))
	}

}
