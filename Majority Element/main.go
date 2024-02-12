package main

import (
	"fmt"
	"sort"
)

func main() {
	result := majorityElement([]int{3, 2, 3})
	fmt.Print(result)
}

type kv struct {
	Value int
	Count int
}

func majorityElement(nums []int) int {

	var ss []kv
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ss = append(ss, kv{nums[0], 0})
	indexResponse := 0
	kvIndex := 0
	for i := 1; i < len(nums); i++ {
		if ss[kvIndex].Value == nums[i] {
			ss[kvIndex].Count += 1

			if ss[kvIndex].Count > ss[indexResponse].Count {
				indexResponse = kvIndex
			}
		} else {
			kvIndex++
			ss = append(ss, kv{nums[i], 0})
		}
	}

	return ss[indexResponse].Value
}
