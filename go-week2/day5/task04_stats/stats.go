package task04_stats

func Sum(numbers []int) int {
	var result int
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return float64(Sum(numbers)) / float64(len(numbers))
}
func MinMax(numbers []int) (min int, max int, ok bool) {
	min = 1
	ok = true
	max = 1
	if len(numbers) == 0 {
		return 0, 0, false
	}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
		if len(numbers) == 1 {
			min = numbers[i]
			max = numbers[i]
			break
		}
	}
	if min == 0 && max == 0 {
		ok = false
	}
	return min, max, ok
}
