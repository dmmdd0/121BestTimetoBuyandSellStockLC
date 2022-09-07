package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	prices = []int{2, 1, 2, 1, 0, 1, 2}
	prices = []int{2, 4, 1}
	prices = []int{5, 8, 2, 3, 1}

	fmt.Println(maxProfitOld(prices))
	fmt.Println(maxProfit(prices))
}

func maxProfitOld(prices []int) int {

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
func maxProfit(price []int) int {
	min := price[0] //first vallue
	max := price[0]
	profit := 0
	for i := range price {
		if price[i] > max {
			max = price[i]
		}
		if price[i] < min {
			min = price[i]
			max = min
		}
		if max-min > profit {
			profit = max - min
		}
	}
	return profit
}
