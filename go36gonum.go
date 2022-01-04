package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Scalar: a number
	var a int = 44
	fmt.Println("Simple scalar: ", a)

	// vectors
	// 1 dim
	// row vector / column vector[list of same datatype]
	// method 1
	myvector := []float64{1.2, 3.4, 4.5, 3.5, 4.4}
	fmt.Println("vector: ", myvector)
	fmt.Printf("%T \n", myvector)
	// method 2
	// Syntax: mat.NewVecDense(size, []type(v1,v2))
	myvectorA := mat.NewVecDense(2, []float64{1.2, 3.4})
	myvectorB := mat.NewVecDense(2, []float64{3.2, 4.4})
	fmt.Printf("%T \n", myvectorA)

	// Dot product
	dp := mat.Dot(myvectorA, myvectorB)
	fmt.Println("Dot vectors: ", dp)

	// Sum of vector
	fmt.Println("Sum :", mat.Sum(myvectorA))

	// Max of vector
	fmt.Println("Max :", mat.Max(myvectorA))

	// matrix
	// creating a matrix
	data := []float64{1.7, 3.4, 3.3, 4.5, 3.3, 1.4}
	mymatrix := mat.NewDense(2, 3, data)
	fmt.Println("my matrix: ", mymatrix)
	fmt.Printf("%T \n", mymatrix)
	fm := mat.Formatted(mymatrix, mat.Prefix(" "))
	fmt.Println("fm: \n", fm)
	matPrint(mymatrix)

	// matrix of zeros
	allzeros := mat.NewDense(3, 2, nil)
	matPrint(allzeros)

	// Shape/Dimension
	fmt.Println(mymatrix.Dims())
	// Get Min/Max
	fmt.Println("min", mat.Min(mymatrix))
	fmt.Println("max", mat.Max(mymatrix))

	// Selection
	// arr[0] : python
	// in GO will not work
	// fmt.Println(mymatrix[0])
	// Get all values of row 1
	row1 := mat.Row(nil, 0, mymatrix)
	fmt.Println(row1)
	// Method 2 : Get all values of a row
	fmt.Println(mymatrix.RowView(0))

	// Get all values of col1
	col1 := mat.Col(nil, 0, mymatrix)
	fmt.Println(col1)

	// method 2: Get all values of a column
	fmt.Println(mymatrix.ColView(0))

	// Select a single value
	fmt.Println(mymatrix.At(0, 0)) // At(Row,Col)

	// Modify/Change
	// arr[0][0] = 34 : numpy python
	// SetCol.Set.SetRow()
	mymatrix.Set(0, 0, 7.1)
	matPrint(mymatrix)

	// Change an entire row
	mymatrix.SetRow(0, []float64{3.0, 3.0, 3.0})
	matPrint(mymatrix)

	// Change an entire collumn
	mymatrix.SetCol(0, []float64{0.0, 0.0})
	matPrint(mymatrix)

	matPrint(mymatrix.T())
}

// to display matrix
func matPrint(X mat.Matrix) {
	fm := mat.Formatted(X, mat.Prefix(" "), mat.Squeeze())
	fmt.Println("%v \n", fm)
}
