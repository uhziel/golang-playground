package main

import "fmt"

func main() {
	tests := []int{
		0,
		1,
		2,
		3,
		4,
		11,
	}
	for _, n := range tests {
		fmt.Println(n, circelReturnStart(n))
		fmt.Println(n, circelReturnStart2(n), "method2")
	}
}

const circelLen = 10

func circelReturnStart(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, circelLen)
	dp[0] = 1

	for i := 0; i < n; i++ {
		dpTmp := make([]int, circelLen)
		for j := 0; j < circelLen; j++ {
			l := j - 1
			if l < 0 {
				l = circelLen - 1
			}
			dpTmp[l] += dp[j]

			r := j + 1
			if r > circelLen-1 {
				r = 0
			}
			dpTmp[r] += dp[j]
		}
		dp = dpTmp
		fmt.Println(dp)
	}

	return dp[0]
}

func circelReturnStart2(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	dp[0] = make([]int, circelLen)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, circelLen)
		for j := 0; j < circelLen; j++ {
			dp[i][j] = dp[i-1][(j-1+circelLen)%circelLen] + dp[i-1][(j+1)%circelLen]
		}
		fmt.Println(dp[i])
	}

	return dp[n][0]
}
