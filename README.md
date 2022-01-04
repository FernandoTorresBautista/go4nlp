# go4nlp
Go For Natural Language Proccessing Resource

- Go For NLP
+ Strings
+ Regex / regexp
+ Language detection 
+ NLP 
  + Tokenization 
  + Pos
  + NER 
  + etc 

+ Sentiment Analysis
+ Text classification 
+ Summary 

### Go 4 NLP Pkgs 
+ Prose: github.com/jdkato/prose/v2
+ NLP: github.com/james-bowman/nlp 
+ Spacy go: github.com/yash1994/spacy-go
+ Ling: github.com/liuzl/ling 
+ Vader Go:
  + github.com/grassmudhorses/vader-go/sentitext
  + "github.com/cdipaolo/sentiment"
+ Regex
  + CommonRegex: go get github.com/mingrammer/commonregex
  + Norm: unicode normalization 


### GO sentiment analysis excersice with Vader 
+ github.com/grassmudhorses/vader-go
  + Vader to DO sentiment analysis on already labelled dataset 
+ gopkg.in/kniren/gota.v0/dataframe 
+ gopkg.in/kniren/gota.v0 [all]


### Go for statictics 
+ math
+ Gonum/stats : gonum.org/v1/gonum/stat
+ Stats: go get github.com/montanaflynn/stats

### Fo for exploratory data analysis 
+ GoTa: go get github.com/kniren/gota/dataframe
+ qframes: github.com/tobgu/qframe
+ dataframes-go 
+ etc 

#### F Structure 
  ```
  type F struct {
    Colidx     int
    Colname    string
    Comparator series.Comparator
    Comparando interface{}
	}
  ```
