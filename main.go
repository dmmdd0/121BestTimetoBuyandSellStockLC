package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	type stock struct {
		day   int
		price int
	}
	var min stock
	var max stock

	for i, p := range prices {
		if p < min.price || min.price == 0 {
			min.day = i
			min.price = p
		}
		if p > max.price && i > min.day {
			max.day = i
			max.price = p
		}

	}
	profit := max.price - min.price
	if profit < 0 {
		profit = 0
	}
	return profit
}
