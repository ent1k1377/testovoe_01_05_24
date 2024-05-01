package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 45; i++ {
		fmt.Printf("%d %s\n", i, getNumberPC(i))
	}
}

func getNumberPC(n int) string {
	fNum := int64(math.Abs(float64(n)))

	if fNum >= 5 && fNum <= 20 {
		return "компьютеров"
	}

	fNum = fNum % 10

	if fNum == 1 {
		return "компьютер"
	} else if fNum > 1 && fNum < 5 {
		return "компьютера"
	} else {
		return "компьютеров"
	}
}
