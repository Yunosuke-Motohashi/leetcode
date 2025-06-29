// import "fmt"
func isSubsequence(s string, t string) bool {
    m, n := len(s), len(t)
    if m > n {
        return false
    }
    dp := make([]bool, n+1)
    for j:=0; j<=n; j++ {
        dp[j] = true
    }
    for i := 1; i<=m; i++ {
        pre := dp[0]
        for j:=1; j<=n; j++ {
            tmp := dp[j]
            dp[0] = false
            if s[i-1] == t[j-1] {
                dp[j] = pre
            } else {
                dp[j] = dp[j-1]
            }
            // fmt.Println(i,j, dp, pre,tmp, s[i-1], t[j-1])
            pre = tmp
        }
    }
    return dp[n]
}