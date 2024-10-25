package main

import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	i, j := 0, 1
	for j < len(nums) {
		if nums[j]-nums[i] > ans {
			ans = nums[j] - nums[i]
		}
		j++
		i++
	}
	return ans
}

func main() {
	nums := []int{3, 6, 9, 1}
	res := maximumGap(nums)
	fmt.Print(res)

}
