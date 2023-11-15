package utils

import (
	"strconv"
	"strings"
)

func DotString2Arr(s string) []int {
	intSlice := make([]int, 0)
	for _, s := range strings.Split(s, ",") {
		if num, err := strconv.Atoi(s); err == nil {
			intSlice = append(intSlice, num)
		}
	}
	return intSlice
}
