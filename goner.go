package main

import (
	"fmt"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {
	// NER
	// Entity: Person/People/Org/Location/etc
	var mytext string = "Fernando Torres works in México as Golang Developer"

	// NLP Document Struct/Object
	doc, err := prose.NewDocument(mytext)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// for index, entity := range doc.Entities() {
	// 	fmt.Println(index, entity.Text, entity.Label)
	// }

	// for _, entity := range doc.Entities() {
	// 	fmt.Println(entity.Text, entity.Label)
	// }

	// Tokenization: Find the entities alongside
	for _, token := range doc.Tokens() {
		fmt.Println(token.Text, token.Label)
	}
}
