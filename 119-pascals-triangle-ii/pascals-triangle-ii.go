import "fmt"
func getRow(rowIndex int) []int {
    dp := make([][]int, rowIndex+1)
    if rowIndex == 0 {
        return []int{1}
    }
    for i := range dp {
        dp[i] = make([]int, i+1)
    }
    dp[0][0] = 1
    for i := 1; i<=rowIndex; i++ {
        for j := 0; j <= i; j++ {
            fmt.Println(i,j)
            if j==0 || j==i {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }
    }
    return dp[rowIndex]
}