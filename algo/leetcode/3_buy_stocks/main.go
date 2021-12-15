package main

func stockProfit(price []int) int {
	//  price: []int{7, 2, 1, 5, 3, 6, 4},
	//	want:  5,
	profit := 0
	for i := 0; i < len(price)-1; i++ {
		for j := i + 1; j < len(price); j++ {
			if price[j] > price[i] {
				dayProfit := price[j] - price[i]
				if dayProfit > profit {
					profit = dayProfit
				}
			}
		}
	}
	return profit
}

func stockProfitOptimized(price []int) int {
	profit := 0
	if len(price) == 0 || price == nil {
		return profit
	}
	for left, right := 0, 1; right < len(price); {
		if price[left] > price[right] && price[right] > 0 {
			left++
			right++
		} else {
			dayProfit := price[right] - price[left]
			if dayProfit > profit {
				profit = dayProfit
			}
			right++
		}
	}
	return profit
}
