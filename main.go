package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	prices = []int{2, 4, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	type stock int

	var min, max, profit int
	var day int

	for i, p := range prices {
		if (p < min && i != len(prices)) || min == 0 {
			day = i
			min = p
		}
		if p > max && i > day {
			max = p
		}

		if profit < max-min {
			profit = max - min
		}

	}
	if profit < 0 {
		profit = 0
	}
	return profit
}
