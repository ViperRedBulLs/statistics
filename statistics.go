package statistics

import (
	"errors"
	"math"
)

// The package statistics

type Array struct {
	Data []float64
}

// Sort is a function to sorting arrays from the lowest value to highest value.
// If reverse argument is set to true so the arrays to sorting from the lowest value
// to the highest value
func (a Array) Sort(reverse bool) {
	n := len(a.Data) // Length of array

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			idx := i
			if reverse && a.Data[j] > a.Data[idx] {
				idx = j
			}
			if !reverse && a.Data[j] < a.Data[idx] {
				idx = j
			}
			temp := a.Data[i]
			a.Data[i] = a.Data[idx]
			a.Data[idx] = temp
		}
	}
}

// Sum is a function to get summary of all the arrays
func (a Array) Sum() float64 {
	var result float64
	for _, value := range a.Data {
		result += value
	}
	return result
}

// Min is a function to get the smallest value of arrays
func (a Array) Min() float64 {
	a.Sort(false)
	return a.Data[0]
}

// Max is a function to get the highest value of arrays
func (a Array) Max() float64 {
	a.Sort(true)
	return a.Data[0]
}

// Median is a function to get the middle value of arrays
func (a Array) Median() float64 {
	a.Sort(false)
	n := len(a.Data)
	half := int(math.Floor(float64(n) / 2))

	if n%2 == 0 {
		return (a.Data[half] + a.Data[half-1]) / 2.0
	} else {
		return a.Data[half]
	}
}

// Range is a function to calculate the difference between the highest
// and smallest values
func (a Array) Range() float64 {
	min := a.Min()
	max := a.Max()
	return max - min
}

// Mods is a function to find the how many data appears
// Using algorithm very high professional
func (a Array) Mods() (float64, error) {
	n := len(a.Data)
	var k float64 = 1
	var x int = 0
	var z int = 0
	manyOfData := make([]float64, n)
	m := make([]float64, n)

	// Count many of data summoning
	for i := 0; i < n; i++ {
		manyOfData[i] = 0
		for j := 0; j < n; j++ {
			if a.Data[i] == a.Data[j] {
				manyOfData[i]++
			}
		}
	}

	// Count how many data appear frequently
	for i := 0; i < n; i++ {
		if manyOfData[i] > k {
			k = manyOfData[i]
		}
	}

	// If mods less than one(1)
	for i := 0; i < n; i++ {
		if x == 0 {
			m[x] = 0
		} else {
			m[x] = m[x-1]
		}

		if manyOfData[i] == k {
			if a.Data[i] != m[x] {
				m[x] = a.Data[i]
				x++
			}
		}
	}

	// If all data appears equal
	for i := 0; i < n; i++ {
		if manyOfData[i] == k {
			z++
		}
	}
	if z == n {
		x = 0
	}
	if x == 0 {
		return 0, errors.New("data not have a mods value")
	} else {
		return m[x], nil
	}
}
