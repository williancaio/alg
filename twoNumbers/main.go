package main

import (
	"fmt"
	"sort"
)

func main() {

}

func twoSum(nums []int, target int) []int {
	sum := 0
	var index1, index2 int = 0, len(nums) - 1

	for i := 0; i < len(nums); i++ {
		sum = nums[index1] + nums[index2]

		if sum == target {
			return []int{index1, index2}
		} else if sum > target {
			index2--
		} else {
			index1++
		}
	}

	return []int{}
}

func Sort(nums []int) [][]int {
	var sortNums [][]int
	for index, line := range nums {
		sortNums[index][index] = line
	}

	sort.Slice(sortNums, func(i, j int) bool {
		return sortNums[i][i] < sortNums[j][0]
	})
	for _, linha := range sortNums {
		fmt.Println(linha)
	}
	return sortNums
}
