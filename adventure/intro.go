package adventure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Intro(p Part, json []byte) string {
	var arc string

	// Prints every line of the story
	for i := 0; i < len(p.Intro.Story); i++ {
		fmt.Println(p.Intro.Story[i])
	}
	fmt.Println()

	// Prints every text option
	for j := 0; j < len(p.Intro.Options); j++ {
		fmt.Printf("Option %d: %s\n", j+1, p.Intro.Options[j].Text)
	}

	fmt.Println("\nPress 1 to continue with Option 1 or 2 to continue with Option 2\n")

	// takes input from the user and throws an error if the input is not 1 or 2
	scanner := bufio.NewReader(os.Stdin)
	option, err := scanner.ReadString('\n')
	option = strings.Trim(option, "\n")

	if err != nil {
		fmt.Println(err)
	} else if option < "1" || option > "2" {
		fmt.Println("\nValid Options are 1 to continue with Option 1 or 2 to continue with Option 2")
	}

	// Sets arc to the corresponding next story arc, depending on the option the user chose
	if option == "1" {
		arc = p.Intro.Options[0].Arc
	} else if option == "2" {
		arc = p.Intro.Options[1].Arc
	}

	fmt.Println(arc)
	return arc
}
