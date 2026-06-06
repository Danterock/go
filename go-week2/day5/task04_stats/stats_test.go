package task04_stats

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if result != 55 {
		t.Error("Expected", 55, "got", result)
	}
}

func TestAverage(t *testing.T) {
	result := Average([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if result != 5.5 {
		t.Error("Expected", 5.5, "got", result)
	}
}
func TestMinMax(t *testing.T) {
	var a, m, ok = MinMax([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if a != 1 {
		t.Error("Expected", 1, "got", a)
	}
	if m != 10 {
		t.Error("Expected", 10, "got", m)
	}
	if !ok {
		t.Error("Expected", true, "got", false)
	}
}
func TestSize(t *testing.T) {
	result := Sum([]int{})
	if result != 0 {
		t.Error("Expected", 0, "got", result)
	}

	res := Average([]int{})

	if res != 0 {
		t.Error("Expected", 0, "got", res)
	}

	_, _, ok := MinMax([]int{})

	if ok {
		t.Error("Expected", false, "got", ok)
	}

}
func TestNegativ(t *testing.T) {
	numbers := []int{-5}

	result := Sum(numbers)
	if result != -5 {
		t.Error("Expected", -5, "got", result)
	}
	res := Average(numbers)
	if res != -5 {
		t.Error("Expected", -5, "got", res)
	}
	min, max, ok := MinMax(numbers)

	if !ok {
		t.Error("Expected", true, "got", false)
	}
	if min != -5 {
		t.Error("Expected", -5, "got", min)
	}
	if max != -5 {
		t.Error("Expected", -5, "got", max)
	}
}
