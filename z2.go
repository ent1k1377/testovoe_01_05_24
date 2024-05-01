package main

import (
	"fmt"
)

func main() {
	q := []int{42, 12, 18}
	fmt.Println(getCommonDivisors(q))
}

func getCommonDivisors(s []int) (res []int) {
	maxNum, err := getMaxSliceNumber(s)
	if err != nil {
		return res
	}

	for i := 2; i < maxNum; i++ {
		if isDivisor(s, i) {
			res = append(res, i)
		}
	}

	return res
}

func isDivisor(s []int, divisor int) bool {
	for _, v := range s {
		if v%divisor != 0 {
			return false
		}
	}

	return true
}

func getMaxSliceNumber(s []int) (int, error) {
	if len(s) == 0 {
		return -1, fmt.Errorf("slice length equal 0")
	}

	res := s[0]
	for i := 1; i < len(s); i++ {
		if res < s[i] {
			res = s[i]
		}
	}

	return res, nil
}
