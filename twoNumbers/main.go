package main

import (
	"fmt"
	"sort"
)

func main() {

	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}

type kv struct {
	Index int
	Value int
}

func twoSum(nums []int, target int) []int {
	sum := 0
	var index1, index2 int = 0, len(nums) - 1
	sortNumbers := Sort(nums)

	for i := 0; i < len(sortNumbers); i++ {
		sum = sortNumbers[index1].Value + sortNumbers[index2].Value

		if sum == target {
			return []int{sortNumbers[index1].Index, sortNumbers[index2].Index}
		} else if sum > target {
			index2--
		} else {
			index1++
		}
	}

	return []int{}
}

func Sort(nums []int) []kv {

	var ss []kv
	for k, v := range nums {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	return ss
}
