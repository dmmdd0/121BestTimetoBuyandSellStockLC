package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
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
		if p < min.price {
			min.day = i
			min.price = p
		}
		if p > max.price {
			max.day = i
			max.price = p
		}

	}
	return 0
}
