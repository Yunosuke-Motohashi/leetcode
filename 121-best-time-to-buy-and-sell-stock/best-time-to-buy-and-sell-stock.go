import "fmt"
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    n := len(prices)
    dp := make([][2]int, n)
    
    // 初期条件
    dp[0][0] = 0           // sell
    dp[0][1] = prices[0]  // buy

    for i := 1; i < n; i++ {
        // sell
        dp[i][0] = max(dp[i-1][0], prices[i]-dp[i-1][1])
        // buy
        dp[i][1] = min(dp[i-1][1], prices[i])
        fmt.Println(dp[i], dp[i-1][0], dp[i-1][1])
    }

    // 最終的に株を持っていない状態の最大利益
    if dp[n-1][0] < 0 {
        return 0
    }
    fmt.Println(dp)
    return dp[n-1][0]
}
