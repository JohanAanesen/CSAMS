package util_test

import (
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
	"math"
	"testing"
)

func TestStatistics(t *testing.T) {
	const MIN = 1
	const MAX = 5
	const SIZE = 30
	const AVG = 3.233333333333333
	const VAR = 1.378888888888888
	const PERCENT = 64.66666666666
	var StdDev = math.Sqrt(VAR)

	var fakeData = []float64{3, 4, 3, 2, 5, 4, 3, 2, 1, 4, 3, 2, 4, 3, 5, 2, 5, 3, 2, 1, 5, 4, 3, 4, 4, 3, 2, 4, 2, 5}

	stats := util.NewStatistics(MIN, MAX)
	stats.Entries = fakeData

	if stats.Size() != SIZE {
		t.Fail()
	}

	avg, err := stats.Average()
	if err != nil {
		t.Fail()
	}

	const ExpectedAvgDiff = 3e-15
	if diff := math.Abs(avg - AVG); diff > ExpectedAvgDiff {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedAvgDiff, diff)
		t.Fail()
	}

	Var, err := stats.Variance()
	if err != nil {
		t.Fail()
	}

	const ExpectedVarDiff = 1e-9
	if diff := math.Abs(Var - VAR); diff > ExpectedVarDiff {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedVarDiff, diff)
		t.Fail()
	}

	stdDev, err := stats.StandardDeviation()
	if err != nil {
		t.Fail()
	}

	const ExpectedStdDevDiff = 1e-9
	if diff := math.Abs(stdDev - StdDev); diff > ExpectedStdDevDiff {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedStdDevDiff, diff)
		t.Fail()
	}

	percent, err := stats.AveragePercent()
	if err != nil {
		t.Fail()
	}

	const ExpectedPercentDiff = 1e-9
	if diff := math.Abs(percent - PERCENT); diff > ExpectedPercentDiff {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedPercentDiff, diff)
		t.Fail()
	}

	const ExpectedMin = 1
	const ExpectedMax = 5

	if min := stats.Min(); ExpectedMin != min {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedMin, min)
		t.Fail()
	}

	if max := stats.Max(); ExpectedMax != max {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", ExpectedMax, max)
		t.Fail()
	}

	if MIN != stats.AbsMin {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", MIN, stats.AbsMin)
		t.Fail()
	}

	if MAX != stats.AbsMax {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", MAX, stats.AbsMax)
		t.Fail()
	}

	var prevLen = len(stats.Entries)
	stats.AddEntry(4)

	if prevLen >= stats.Size() {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", prevLen+1, prevLen)
		t.Fail()
	}

	// Testing zero division
	stats.Entries = make([]float64, 0)
	stats.AbsMax = 0

	_, err = stats.Average()
	if err == nil {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", util.ErrDivideByZero, nil)
		t.Fail()
	}

	_, err = stats.AveragePercent()
	if err == nil {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", util.ErrDivideByZero, nil)
		t.Fail()
	}

	_, err = stats.Variance()
	if err == nil {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", util.ErrDivideByZero, nil)
		t.Fail()
	}

	_, err = stats.StandardDeviation()
	if err == nil {
		t.Logf("\nexpected:\t%v\ngot:\t\t%v", util.ErrDivideByZero, nil)
		t.Fail()
	}
}

func BenchmarkStatistics(b *testing.B) {

}
