package main

import (
	"fmt"
	"log"

	"github.com/ViperRedBulLs/statistics"
)

func main() {
	a := statistics.Array{
		Data: []float64{1, 2, 3, 4, 5, 6, 7, 3},
	}

	fmt.Println("Real Data\t\t:\t", a.Data)

	// Sorting array
	a.Sort(false)
	fmt.Println("Sorting Data\t\t:\t", a.Data)

	// Sum
	fmt.Println("Summary Data\t\t:\t", a.Sum())

	// Min
	fmt.Println("Min Data\t\t:\t", a.Min())

	// Max
	fmt.Println("Max Data\t\t:\t", a.Max())

	// Median
	fmt.Println("Median Data\t\t:\t", a.Median())

	// Range
	fmt.Println("Range Data\t\t:\t", a.Range())

	mods, err := a.Mods()
	log.SetPrefix("Mods Data\t\t:\t")
	log.SetFlags(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mods Data\t\t:\t", mods)
}
