package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jdkato/prose/tokenize"
)

func main() {

	// Tokenization :split or segment sentence into tokens/units
	// From Scratch
	var mytext string = "Paul wasn't coding at all"
	// Method 1
	tokens := strings.Split(mytext, " ")
	fmt.Println("Text:", mytext)
	fmt.Println(tokens)

	// Method 2: Rule Based (Regex)
	r := regexp.MustCompile(`\w+`)
	newtokens := r.FindAllString(mytext, -1)
	fmt.Println(newtokens)

	// Method 3: Regex + Split
	r2 := regexp.MustCompile(" ")
	newtokens2 := r2.Split(mytext, -1)
	fmt.Println(newtokens2)

	// Packages

	tokenizer := tokenize.NewTreebankWordTokenizer()
	for _, tok := range tokenizer.Tokenize(mytext) {
		fmt.Println(tok)
	}
	wordtokens := tokenizer.Tokenize(mytext)
	fmt.Println(wordtokens)
}
