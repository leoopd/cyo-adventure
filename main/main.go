package main

import (
	utils "github.com/leoopd/cyo-adventure/utils"
)

func main() {

	// defer os.Exit(1)

	file := "gopher.json"
	utils.StoryTeller(utils.JSONParser(file))

}
