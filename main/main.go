package main

import "github.com/leoopd/cyo-adventure/adventure"

func main() {

	file := "gopher.json"
	var advPart adventure.Part

	_ = adventure.Intro(advPart, adventure.JSONParser(file))

}
