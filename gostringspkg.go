package main

import (
	"fmt"
	"strings"
)

func main() {
	// working with strings & string package
	// a string: collection[array/slice] or bytes or runes which characters
	var mystr string = "go String"
	fmt.Println(mystr)
	//check for the type
	fmt.Printf("%T \n", mystr)

	// string manipulation
	//length of string
	fmt.Println("Length: ", len(mystr))
	fmt.Println("Uppercase: ", strings.ToUpper(mystr)) // to uppercase
	fmt.Println("Lowercase: ", strings.ToLower(mystr)) // to lowercase
	fmt.Println("Titlecase: ", strings.ToTitle(mystr)) // to titlecas
	// to get the count of subset within our string
	fmt.Println("TCount: ", strings.Count(mystr, "g")) // how mani times appear g
	// check if contains something
	fmt.Println("Contains: ", strings.Contains(mystr, "o "))

	// Useful for NLP
	// Split: Tokenazation
	tokens := strings.Split(mystr, " ")
	fmt.Println(tokens)
	fmt.Println(tokens[0])
	// Split After
	ex := "go,4,data,science"
	fmt.Println(strings.SplitAfter(ex, "4"))

	// Replace: text clean
	fmt.Println(strings.Replace("go:go strings", "go", "Golang", 1))
	fmt.Println(strings.Replace("go:go strings", "go", "Golang", -1))

	// Join
	f1 := []string{"N", "L", "P"}
	fmt.Println(strings.Join(f1, ""))

	// Subset/Indexing: Stemming
	system := "running"
	fmt.Println(len(system), "[0:3]", system[0:3])

	// Prefix/Suffix
	fmt.Println("Prefix:ru ", strings.HasPrefix(system, "ru"))
	fmt.Println("Suffix:ing ", strings.HasSuffix(system, "ing"))
}
