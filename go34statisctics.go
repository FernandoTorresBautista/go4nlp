package main

import (
	"fmt"
	"math"

	"github.com/montanaflynn/stats"
	gstat "gonum.org/v1/gonum/stat"
)

func main() {
	fmt.Println("Statistics in Go")

	// array / slice
	even := []float64{2, 4, 6, 8, 10, 8, 8}
	var odd = []float64{3, 5, 7, 9}
	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)

	// basic maths
	// mean
	evenMean, _ := stats.Mean(even)
	// if err != nil {
	// 	log.Println(err)
	// }
	fmt.Println("even mean: ", evenMean)
	evenMin, _ := stats.Min(even)
	fmt.Println("even minimum: ", evenMin)
	evenMax, _ := stats.Max(even)
	fmt.Println("even maximum: ", evenMax)

	// Mode
	evenmode, _ := stats.Mode(even)
	fmt.Println("even Mode: ", evenmode)

	// Arithmetic, Harm, Geo
	gmean, _ := stats.GeometricMean(even)
	fmt.Println("GeoMean: ", gmean)

	harmean, _ := stats.HarmonicMean(even)
	fmt.Println("HarmonicMean: ", harmean)

	// STD
	oddstd, _ := stats.StandardDeviation(odd)
	fmt.Println("Standard deviation: ", oddstd)

	// Variance
	oddvar, _ := stats.PopulationVariance(odd)
	fmt.Println("Population Variance: ", oddvar)

	// Math Packages
	fmt.Println("SQRT DEV from math: ", math.Sqrt(oddvar))

	// USING GONUM For Stats
	//
	oddmean := gstat.Mean(odd, nil)
	fmt.Println("Mean for Odd via Gonum: ", oddmean)
}
