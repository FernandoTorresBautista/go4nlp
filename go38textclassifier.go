package main

import (
	"fmt"
	"log"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/text"
)

func main() {
	// Use CSV
	// Normal approach
	// Stream of data

	// Chan: channel to communicate between go rutines
	stream := make(chan base.TextDatapoint, 100)
	// Error
	errors := make(chan error)

	// Init model
	// (stream <-chan base.TextDatapoint, classes uint8, sanitize func(rune) bool) *text.NaiveBayes
	model := text.NewNaiveBayes(stream, 2, base.OnlyWordsAndNumbers)

	// Training a model
	go model.OnlineLearn(errors)

	// Classify a text as either hardware or software
	// Parse our data
	stream <- base.TextDatapoint{
		X: "He bought a laptop", // features
		Y: 0,                    // label
	}
	stream <- base.TextDatapoint{
		X: "my unit test is failed",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "tried the program, but is was buggy",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "i need a new power supply",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "the drive has a 2TB capacity",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "unit-tests",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "program",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "power supply",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "drive",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "it need more memory",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "code",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "i found some bugs",
		Y: 0,
	}
	stream <- base.TextDatapoint{
		X: "i swapped the memory",
		Y: 1,
	}
	stream <- base.TextDatapoint{
		X: "i tested the code",
		Y: 0,
	}

	// Close
	close(stream)

	for {
		err, more := <-errors
		if err != nil {
			log.Fatal(err)
		} else {
			break
		}
		fmt.Println(more)
	}

	// Labels: Hardware 1, Software 0
	ex1 := "John bought a laptop"          // Hardware 1
	ex2 := "He fixed the bugs in the code" // Software 0

	mypred1 := model.Predict(ex1)
	mypred2 := model.Predict(ex2)
	fmt.Println(ex1, ", Label: ", mypred1, "\n")
	fmt.Println(ex2, ", Label: ", mypred2)

	// save model
	model.PersistToFile("data/hardvssoftwareclassfier.json")
}
