// import "fmt"

func tribonacci(n int) int {
	dp := []int{0, 1, 1}
	tribo := 1
	for i := 3; i <= n; i++ {
		tribo = tribo + dp[0]
		dp = append(dp, dp[i-1] + dp[i-2] + dp[i-3])
		// fmt.Println(dp, tribo)
	}
	return dp[n]
}