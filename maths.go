package sensors

import (
	"sort"
)

// average sums all values and divides over the number of items in the slice
func average(in []float64) float64 {
	n := float64(len(in))
	if n == 0 {
		return 0
	}
	sum := float64(0)
	for _, x := range in {
		sum += x
	}

	return round2decimals(sum / n)
}

// median gets the value in the middle of the array
// for even number of values it will return the lesser value between the two middle ones
func median(in []float64) float64 {
	n := len(in)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return in[0]
	}
	sort.Float64s(in)
	return in[n/2]
}

//
func mode(in []float64) []float64 {
	result := []float64{}
	n := len(in)
	if n == 0 {
		return result
	}

	sort.Float64s(in)

	last := float64(0)
	count := 1
	// holds how many times a value is found e.g. [1][1,2,3],[2][2,3],[3][3]
	// 3 appears 3 times, if we get the maxmimum value of this map, then we will have the answer to mode()
	// if we had  [1][1,2,3],[2][2,3], then the answe would be [2,3]
	t := map[int][]float64{}
	for k, x := range in {
		if last == x && k != 0 { // skip first index if x == 0
			count++
		} else { // keeps count of how many times x has appeared in the slice
			last = x
			count = 1
		}

		t[count] = append(t[count], x)
	}
	indices := []int{}
	for k := range t {
		indices = append(indices, k)
	}
	if len(t) == 1 { // no value appears more than once so the mode should be empty
		return []float64{}
	}

	return t[max(indices)]
}

// max sorts the input and returns the last value in the slice
func max(in []int) int {
	if len(in) == 0 {
		return 0
	}

	sort.Ints(in)

	return in[len(in)-1]
}
