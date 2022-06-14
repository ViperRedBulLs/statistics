package statistics_test

import (
	"testing"

	"github.com/ViperRedBulLs/statistics"
)

var data = []float64{4, 3, 2, 1, 6, 8, 9, 0}

var (
	dataMedian float64 = 3.5
	dataMin    float64 = 0
	dataMax    float64 = 9
	dataSum    float64 = 33
	dataAvg    float64 = 4.125
)

// Testing Median
func TestMedianArray(t *testing.T) {
	t.Log("Testing Median")

	if statistics.Median(data) != dataMedian {
		t.Errorf("Error Median of `%v` should `%v`.", data, dataMedian)
	}
}

func TestMinArray(t *testing.T) {
	t.Logf("Testing Min")
	if statistics.Min(data) != dataMin {
		t.Errorf("Error `Min` of `%v` should `%v`", data, dataMin)
	}
}

func TestMaxArray(t *testing.T) {
	t.Log("Testing Max")
	if statistics.Max(data) != dataMax {
		t.Errorf("Error `Max` of `%v` should `%v`", data, dataMax)
	}
}

// Sum
func TestSum(t *testing.T) {
	t.Log("Testing Sum")

	if statistics.Sum(data) != dataSum {
		t.Errorf("Error `SUM` of `%v` should `%v`", data, dataSum)
	}
}

// Average
func TestAverage(t *testing.T) {
	t.Log("Testing Average")

	if statistics.Average(data) != dataAvg {
		t.Errorf("Error `AVERAGE` of `%v` should `%v`", data, dataSum)
	}
}
