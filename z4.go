package main

import (
	"fmt"
	"strconv"
)

func main() {
	printMultiplicationTable(5)
}

func printMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("\t%d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Print(i)
		for j := 1; j <= n; j++ {
			fmt.Printf("\t%s ", strconv.Itoa(i*j))
		}
		fmt.Println()
	}
}
