package main

import (
	"fmt"

	"github.com/rylans/getlang"
)

func main() {
	var mystr string = `hello world everyone`
	var mystr2 string = `Bonjour a tous`

	langInfo := getlang.FromString(mystr)
	langInfo2 := getlang.FromString(mystr2)

	fmt.Println("Text: ", mystr)
	fmt.Println(langInfo.LanguageCode()) // language as code
	fmt.Println(langInfo.Confidence())   // confidence/accurrancy

	fmt.Println("Text 2: ", mystr2)
	fmt.Println(langInfo2.LanguageCode()) // language as code
	fmt.Println(langInfo2.Confidence())   // confidence/accurrancy
}
