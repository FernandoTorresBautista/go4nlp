package main

import "fmt"

func main() {
	fmt.Println("Hello Go Lovers")

	// Strings
	// string : a collection of characters
	// string: an array/slice of bytes/runes

	var char byte = 'A'
	var char2 rune = 'A'
	fmt.Println("char byte: ", char, "\nchar rune: ", char2)
	// Types
	fmt.Printf("%T \n", char)  // byte uint8
	fmt.Printf("%T \n", char2) // rune int32

	// method2: create characters with method
	chara := byte('A')
	charb := rune('A')
	fmt.Println("char byte:fxn: ", chara, "\nchar rune:fxn: ", charb)

	// Actual representation
	str1 := string(char)
	str2 := string(char2)
	fmt.Println((str1))
	fmt.Println((str2))
	fmt.Printf("%T \n", str1)
	fmt.Printf("%T \n", str2)

	// representing :method2 %c + PrintF
	fmt.Printf("%c", char)
	fmt.Printf("%c", char2)

}
