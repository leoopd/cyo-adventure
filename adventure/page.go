package adventure

// import (
// 	"html/template"
// 	"log"
// 	"os"
// 	"text/template"

// 	"github.com/leoopd/cyo-adventure/adventure"
// )

// func adventureSite(a adventure.AdvPart) {
// 	// t1 := template.New("Title")
// 	// t1, err := t1.Parse("{{.Title}}\n")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// t1.Execute(os.Stdout, a)

// 	// t2 := template.New("Story")
// 	// t2, err = t2.Parse("{{.Story}}\n")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// t2.Execute(os.Stdout, a)

// 	// t3 := template.New("Options")
// 	// t3, err = t3.Parse("{{.Options}}")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// t3.Execute(os.Stdout, a)

// 	t, err := template.New("adventure").ParseFiles("./adventure.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = t.Execute(os.Stdout, a)
// }
