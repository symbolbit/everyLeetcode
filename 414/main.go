package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	hash := make(map[int]bool, len(nums))
	for _, v := range nums {
		hash[v] = true
	}
	var arr []int
	for k, _ := range hash {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	if len(arr) < 3 {
		return arr[len(arr)-1]
	}
	return arr[len(arr)-3]
}

func main() {
	arr := []int{1, 2, 2, 5, 3, 5}
	res := thirdMax(arr)
	fmt.Print(res)
}
