package main

import (
	"fmt"
	"log"
	"math"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func main() {
	fmt.Println("go for machine learning")
	// // Requirement for dataset(with python and pandas)
	// // Numerical
	// // No missing values & headers
	// // Train/test CSV
	// // [xfeature]<target>

	// Load our dataset
	xtrain, ytrain, err := base.LoadDataFromCSV("data/prepdata/trainhcvData.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Training dataset")
	fmt.Printf("%T \n", xtrain)
	fmt.Printf("%T \n", ytrain)

	// Testing data
	xtest, ytest, err := base.LoadDataFromCSV("data/prepdata/testhcvData.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Testing dataset")
	fmt.Printf("%T \n", xtest)
	fmt.Printf("%T \n", ytest)

	// initialize model
	// optimization method()
	// learning rate
	// regularization: for overfitting
	// dataset (Xfeactures)
	// class (binary 0 and 1 )
	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xtrain, ytrain)
	// Train
	err = model.Learn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Training Done")
	fmt.Println("Prediction")
	// Prediction
	s1 := []float64{32.0, 1.0, 38.5, 52.5, 7.7, 22.1, 7.5, 6.93, 3.23, 106.0, 12.1, 69.0}
	// Negative no Hep
	s2 := []float64{62.0, 0.0, 32.0, 416.6, 5.9, 110.3, 50.0, 5.57, 6.3, 55.7, 650.9, 68.5}
	// fmt.Println(s1, s2)
	// Hep
	mypred1, err := model.Predict(s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mypred1)
	fmt.Println("Round Prediction")
	fmt.Println(math.Round(mypred1[0]))

	mypred2, err := model.Predict(s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mypred2)
	fmt.Println("Round Prediction")
	fmt.Println(math.Round(mypred2[0]))

	// // Save model
	// fmt.Println("Save model")
	// model.PersistToFile("logisticHCVmodel.json")

	// Evaluate the model
	fmt.Println("Evaluate Model")
	res := evaluatemodel(xtrain, ytrain, s1, ytest)
	fmt.Printf("%+v \n", res)

}

// confusion matrix struct

type ConfusionMatrix struct {
	positive      int     `json:"positive"`
	negative      int     `json:"positive"`
	truePositive  int     `json:"truePositive"`
	trueNegative  int     `json:"trueNegative"`
	falsePositive int     `json:"falsePositive"`
	falseNegative int     `json:"falseNegative"`
	acurracy      float64 `json:"acurracy"`
	precision     float64 `json:"precision"`
	recall        float64 `json:"recall"`
}

func evaluatemodel(xtrain [][]float64, ytrain []float64, xTest, yTest []float64) ConfusionMatrix {
	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xtrain, ytrain)
	err := model.Learn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finishing Training")

	myprediction, err := model.Predict(xTest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myprediction)
	fmt.Println(math.Round(myprediction[0]))
	cm := ConfusionMatrix{}
	for _, y := range yTest {
		if y == 1.0 {
			cm.positive++
		} else if y == 0.0 {
			cm.negative++
		}
	}

	// Evaluate the model
	var desicionBoundary = 0.5
	for i := range xTest {
		prediction, err := model.Predict(xTest)
		if err != nil {
			panic(err)
		}
		y := int(yTest[i])
		positive := prediction[0] >= desicionBoundary
		if y == 1 && positive {
			cm.truePositive++
		} else if y == 1 && !positive {
			cm.falseNegative++
		} else if y == 0 && positive {
			cm.falsePositive++
		} else if y == 0 && !positive {
			cm.trueNegative++
		}
	}
	fmt.Println("Calculating Confusion Metrics")
	// Calculate Evaluation Metrics
	cm.acurracy = (float64(cm.truePositive) + float64(cm.trueNegative)) /
		(float64(cm.positive) + float64(cm.negative))

	cm.precision = float64(cm.truePositive) / (float64(cm.truePositive) + float64(cm.falsePositive))
	cm.recall = float64(cm.truePositive) / (float64(cm.truePositive) + float64(cm.falseNegative))
	return cm
}
