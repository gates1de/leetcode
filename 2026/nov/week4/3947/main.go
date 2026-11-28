package main

import (
	"fmt"
	"sort"
)

func maximumSaleItems(items [][]int, budget int) int {
	if len(items) == 0 || budget == 0 {
		return 0
	}

	n := len(items)
	frequency := make([]int, n+1)
	minimumPrice := int64(budget) + 1
	for _, item := range items {
		frequency[item[0]]++
		minimumPrice = min(minimumPrice, int64(item[1]))
	}

	divisible := make([]int, n+1)
	for factor := 1; factor <= n; factor++ {
		for multiple := factor; multiple <= n; multiple += factor {
			divisible[factor] += frequency[multiple]
		}
	}

	bonusItems := make([]bonusItem, 0, n)
	for _, item := range items {
		price := int64(item[1])
		bonusCopies := divisible[item[0]] - 1
		if bonusCopies > 0 && price < minimumPrice*2 {
			bonusItems = append(bonusItems, bonusItem{
				price: price,
				count: int64(bonusCopies),
			})
		}
	}
	sort.Slice(bonusItems, func(i int, j int) bool {
		return bonusItems[i].price < bonusItems[j].price
	})

	budgetValue := int64(budget)
	spent := int64(0)
	bonusCount := int64(0)
	for _, item := range bonusItems {
		count := min(item.count, (budgetValue-spent)/item.price)
		spent += count * item.price
		bonusCount += count
		if spent == budgetValue {
			break
		}
	}

	result := 2*bonusCount + (budgetValue-spent)/minimumPrice
	return int(result)
}

type bonusItem struct {
	price int64
	count int64
}

func main() {
	// result: 5
	// items := [][]int{{1,6},{2,4},{3,5}}
	// budget := 19

	// result: 7
	items := [][]int{{2, 8}, {1, 10}, {6, 6}, {4, 12}, {5, 20}, {5, 17}}
	budget := 35

	// result: 0
	// items := [][]int{}
	// budget := 0

	result := maximumSaleItems(items, budget)
	fmt.Printf("result = %v\n", result)
}
