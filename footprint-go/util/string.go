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

func String2Uint(s string) (uint, error) {
	num, err := strconv.ParseUint(s, 10, 0)
	return uint(num), err
}

func String2Int(s string) (int, error) {
	num, err := strconv.ParseUint(s, 10, 0)
	return int(num), err
}
