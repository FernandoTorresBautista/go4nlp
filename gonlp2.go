package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {
	var mystr string = "Hello world this is Golang"
	fmt.Println(mystr)

	// NLP Document Struct
	// docx, _ := prose.NewDocument(mystr)
	// fmt.Println("%T \n", docx)

	// Showing Error
	docx, err := prose.NewDocument(mystr)
	if err != nil {
		log.Fatal("Err ", err)
	}

	// Tokenization
	for i, tok := range docx.Tokens() {
		fmt.Println(i, tok.Text)
	}

	// Reading from a textfile
	content, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal("Err ", err)
	}
	docxfile, err := prose.NewDocument(string(content))
	if err != nil {
		log.Fatal("Err ", err)
	}
	// tokens
	for i, tok := range docxfile.Tokens() {
		fmt.Println(i, tok.Text)
	}
	// sentence
	for i, sent := range docxfile.Sentences() {
		fmt.Println(i, sent.Text)
	}

	// POS Tagging
	for _, tok := range docxfile.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
	}
}
