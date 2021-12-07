package main

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

func main() {
	mytext := "I love eating apples and coding"

	// Model : restore or retrain(project directory)
	sentimentModel, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	results := sentimentModel.SentimentAnalysis(mytext, sentiment.English)
	// 1 positive 0 negative
	fmt.Println(results)
	// Semtiment for the entire sentence
	fmt.Println("Sentiment score: ", results.Score)
}
