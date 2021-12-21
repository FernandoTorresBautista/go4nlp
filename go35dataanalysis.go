package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/gonum/stat"
)

type F struct {
	Colname    string
	Comparator interface{}
	Comparando interface{}
}

func main() {
	// DAta Analysis in Go
	// Open CSV
	//csvfile, err := os.Open("data/diamonds.csv")
	csvfile, err := os.Open("data/amazondataset.csv")
	if err != nil {
		log.Fatal(err)
	}
	// Read CSV
	df := dataframe.ReadCSV(csvfile)
	// fmt.Println(df)
	// EDA

	// Shape of Dataset
	row, col := df.Dims()
	fmt.Println("Shape of DF: ", row, col)

	// Get Only Row Size
	fmt.Println(df.Nrow())

	// Get Only Columns Size
	fmt.Println(df.Ncol())

	// Get Column Names
	fmt.Println("Names: ", df.Names())

	// // Get DataTypes
	// fmt.Println("Types: ", df.Types())

	// // Describe/Summary
	// fmt.Println("Summary: ", df.Describe())

	// // Selection of Columns & Rows
	// // Select columns by Column name
	// fmt.Println("Select column by name: ", df.Select("sentences"))
	// // Select columns by index
	// fmt.Println("Select column by index: ", df.Select(0))

	// Multiple columns selection
	// df[["carat", "cut"]] // pandas
	// []string{"carat", "cut"} // slice / array
	// fmt.Println(df.Select([]string{"carat", "cut"}))
	fmt.Println(df.Select([]string{"sentences", "label"}))

	// Selection of Rows
	// Subset : iloc
	fmt.Println(df.Subset(0))

	// Series and Apply functions
	ds := df.Col("label") // df.Col("carat")
	// fmt.Printf("%T \n", ds)
	// Apply a function
	// Get the Mean
	dsmean := ds.Mean()
	fmt.Println("DS Mean of Series: ", dsmean)

	// Check for missing values
	// fmt.Println(ds.IsNaN())
	gmean := stat.Mean(ds.Float(), nil)
	fmt.Println("Go Num Mean for series: ", gmean)

	// Apply Conditions
	// fmt.Println(ds.Select("cut"))
	// ispremium := df.Filter(dataframe.F{"cut", "==", "1"})
	// fmt.Println(df.Select("label"))
	ispremium := df.Filter(dataframe.F{1, "label", "==", "1"})
	fmt.Println("Is premium: ", ispremium)

}
