package main

import (
	"fmt"
	"io/ioutil"
	"log"

	rake "github.com/Obaied/RAKE.Go"
)

func main() {
	// var mytext string = `
	// Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.
	// Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.
	// A major drawback of statistical methods is that they require elaborate feature engineering. Since the early 2010s,[16] the field has thus largely abandoned statistical methods and shifted to neural networks for machine learning. Popular techniques include the use of word embeddings to capture semantic properties of words, and an increase in end-to-end learning of a higher-level task (e.g., question answering) instead of relying on a pipeline of separate intermediate tasks (e.g., part-of-speech tagging and dependency parsing). In some areas, this shift has entailed substantial changes in how NLP systems are designed, such that deep neural network-based approaches may be viewed as a new paradigm distinct from statistical natural language processing. For instance, the term neural machine translation (NMT) emphasizes the fact that deep learning-based approaches to machine translation directly learn sequence-to-sequence transformations, obviating the need for intermediate steps such as word alignment and language modeling that was used in statistical machine translation (SMT).`

	// reading from a file
	content, err := ioutil.ReadFile("exampletext.txt")
	if err != nil {
		log.Fatal("Error ", err)
	}
	// candidates := rake.RunRake(mytext)
	candidates := rake.RunRake(string(content))

	// Store in a Map [k:v]/Dictionary
	keywordMap := make(map[string]float64, candidates.Len())
	for _, word := range candidates {
		fmt.Println("Keyword: ", word.Key, ", Score: ", word.Value)
		keywordMap[word.Key] = word.Value
	}

	fmt.Println(keywordMap)

}
