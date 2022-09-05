package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	prices = []int{2, 4, 1}
	prices = []int{5, 8, 2, 3, 1}
	prices = []int{2, 1, 2, 1, 0, 1, 2}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	type stock int

	var min, max, profit int
	var day int

	for i, p := range prices {
		if (p < min && i != len(prices)-1) || i == 0 {
			day = i
			min = p
			max = 0
		}
		if p > max && i > day {
			max = p
		}

		if profit < max-min {
			profit = max - min
		}

	}

	return profit
}
