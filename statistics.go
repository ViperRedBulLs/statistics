package statistics

import (
	"math"
)

// Sort is a function to sorting an array
func Sort(data []float64, reverse bool) []float64 {
	n := len(data) // length of array

	// If reverse is not true
	if !reverse {
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				minIdx := i
				if data[j] < data[minIdx] {
					minIdx = j
				}
				temp := data[i]
				data[i] = data[minIdx]
				data[minIdx] = temp
			}
		}
	} else {
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				maxIdx := i
				if data[j] > data[maxIdx] {
					maxIdx = j
				}
				temp := data[i]
				data[i] = data[maxIdx]
				data[maxIdx] = temp
			}
		}
	}

	return data
}

// Median is a function to search a middle value of an array.
func Median(data []float64) (result float64) {
	data = Sort(data, false)
	result = 0
	n := len(data)
	half := int(math.Floor(float64(n) / 2))
	if n%2 == 0 {
		result = (data[half-1] + data[half]) / 2.0
	} else {
		result = data[half]
	}
	return
}

// Min is a function to get the smallest value
func Min(data []float64) (min float64) {
	data = Sort(data, false)
	min = data[0]
	return
}

// Max is a function to get the largest value
func Max(data []float64) (max float64) {
	data = Sort(data, true)
	max = data[0]
	return
}

// Sum is a function that functions to find the sum of all elements
func Sum(data []float64) (result float64) {
	result = 0
	for _, value := range data {
		result += value
	}
	return
}

// Range is a function to Finding the value of
// the difference between the largest value and the smallest value
func Range(data []float64) float64 {
	min := Min(data)
	max := Max(data)
	return max - min
}

// Average is a function to get average of array
func Average(data []float64) float64 {
	n := len(data)
	sum := Sum(data)
	return sum / float64(n)
}
