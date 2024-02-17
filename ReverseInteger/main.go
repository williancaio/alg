package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	myStr := strconv.Itoa(x)
	reverseStr := ""
	for _, v := range myStr {
		reverseStr = string(v) + reverseStr
	}

	if x < 0 {
		reverseStr = strings.ReplaceAll(reverseStr, "-", "")
		reverseStr = "-" + reverseStr
	}
	x, _ = strconv.Atoi(reverseStr)
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}
	return x
}

func main() {
	resp := reverse(1534236469)
	fmt.Println(resp)

	resp = reverse(321)
	fmt.Println(resp)

	resp = reverse(-321)
	fmt.Println(resp)
}
