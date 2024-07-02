package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []float64{2, 1, 3}
	nums := make([]float64, len(numbers))
	copy(nums, numbers)
	sort.Float64s(nums)
	var median float64
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		median = nums[i]
	} else {
		median = (nums[i-1] + nums[i]) / 2
	}
	fmt.Println(numbers)
	fmt.Println(median)

}
