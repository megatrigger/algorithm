package lc

import (
	"fmt"
	"sort"
)

func CoinChange(coins []int, amount int) int {
	sort.Ints(coins)
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	r := -1
	for i := range coins {
		n := 0
		m := amount
		for j := i; j < len(coins); j++ {
			if n += m / coins[j]; r != -1 && n >= r {
				break
			}
			fmt.Println(coins[j], n)
			m = m % coins[j]
			if m == 0 {
				r = n
				fmt.Println("r:", r)
				break
			} else {
				fmt.Println("mod:", m)
			}
		}
		fmt.Println("=======")
	}
	return r
}

func CoinChangeByDp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	for _, v := range coins {
		if v < len(dp) {
			dp[v] = 1
		}
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {// 求构成金额i的最少硬币数
		if dp[i] == 1 {
			continue
		}

		for _, v := range coins {
			if v > i {
				continue
			}
			if dp[i-v] == -1 {
				continue
			}
			n := dp[i-v] + 1
			if dp[i] == -1 || (dp[i] != -1 && n < dp[i]) {
				dp[i] = n
			}
		}
	}
	return dp[len(dp)-1]
}
