package main

import (
	"github.com/leoopd/cyo-adventure/src/utils"
)

func main() {

	// defer os.Exit(1)

	file := "gopher.json"
	utils.JSONParser(file)
}
