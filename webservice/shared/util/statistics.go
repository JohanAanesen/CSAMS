package util

import (
	"errors"
	"math"
)

var (
	// ErrDivideByZero error
	ErrDivideByZero = errors.New("cannot divide by zero")
)

// Statistics struct
type Statistics struct {
	Entries []float64
	AbsMin  float64
	AbsMax  float64
}

// NewStatistics func
func NewStatistics(min, max float64) *Statistics {
	return &Statistics{
		Entries: make([]float64, 0),
		AbsMin:  min,
		AbsMax:  max,
	}
}

// AddEntry adds new entry to the slice of entries
func (s *Statistics) AddEntry(entry float64) {
	s.Entries = append(s.Entries, entry)
}

// Sum up all entries
func (s Statistics) Sum() float64 {
	var sum float64
	for _, entry := range s.Entries {
		sum += entry
	}
	return sum
}

// Size of the entries slice
func (s Statistics) Size() int {
	return len(s.Entries)
}

// Average value of the all the entries
func (s Statistics) Average() (float64, error) {
	if s.Size() == 0 {
		return 0, ErrDivideByZero
	}

	return s.Sum() / float64(s.Size()), nil
}

// AveragePercent of all entries based on a set absolute max
func (s Statistics) AveragePercent() (float64, error) {
	if s.AbsMax == 0 {
		return 0, ErrDivideByZero
	}

	avg, _ := s.Average()

	return (avg / s.AbsMax) * 100, nil
}

// Variance func
func (s Statistics) Variance() (float64, error) {
	avg, err := s.Average()
	if err != nil {
		return 0, err
	}

	var Var float64

	for _, entry := range s.Entries {
		res := entry - avg
		Var += math.Pow(res, 2)
	}

	Var = Var / (float64(s.Size() - 1))

	return Var, nil
}

// StandardDeviation func
func (s Statistics) StandardDeviation() (float64, error) {
	Var, err := s.Variance()
	if err != nil {
		return 0, err
	}

	return math.Sqrt(Var), nil
}

// Min func
func (s Statistics) Min() float64 {
	var min = math.MaxFloat64

	for _, entry := range s.Entries {
		if entry < min {
			min = entry
		}
	}

	return min
}

// Max func
func (s Statistics) Max() float64 {
	var max = -math.MaxFloat64

	for _, entry := range s.Entries {
		if entry > max {
			max = entry
		}
	}

	return max
}
