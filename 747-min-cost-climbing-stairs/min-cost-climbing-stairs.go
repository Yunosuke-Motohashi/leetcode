// import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
    dp[1] = cost[0]
	for i := 2; i < n; i++ {
        dp[i] = cost[i-1] + min(dp[i-2], dp[i-1])
		// fmt.Println(i, dp[i-1], cost[i-1], cost[i], dp)
	}
    dp[n] = min(dp[n-2]+cost[n-1], dp[n-1])
	// fmt.Println(dp)
	return dp[n]
}