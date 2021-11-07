package main

import (
	"fmt"

	"github.com/abadojack/whatlanggo"
)

func main() {

	var mydocx string = "hello world this is Fer"
	lang := whatlanggo.Detect(mydocx)
	fmt.Println(mydocx)
	fmt.Println(lang.Lang.String())
	fmt.Println(lang.Confidence)

	var mydocx2 string = "Hola mundo este es fernando, hablo español"
	lang2 := whatlanggo.Detect(mydocx2)
	fmt.Println(lang2.Lang.String())
	fmt.Println(lang2.Confidence)

	fmt.Println(whatlanggo.Scripts[lang2.Script])
}
