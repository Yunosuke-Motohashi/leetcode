
func generate(numRows int) [][]int {

    dp := make([][]int, numRows)
    for i := range dp {
        dp[i] = make([]int, i+1)
    }
    dp[0][0] = 1
    for i:=1; i<numRows; i++ {
        for j:=0; j<=i; j++ {
            if j == 0 || j == i {
                dp[i][j] = 1
            } else {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            }
        }

    }

    fmt.Println(dp)


    return dp

}