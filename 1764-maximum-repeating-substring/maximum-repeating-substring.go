// import "fmt"
func maxRepeating(sequence string, word string) int {
    m, n := len(sequence), len(word)
    dp := make([]int, m+1)
    ans := 0
    
    for i := m - n; i >= 0; i-- {
        // fmt.Println(i, sequence[i:i+n], dp, ans, sequence[i:i+n] == word)
        if sequence[i:i+n] == word {
            dp[i] = dp[i+n] + 1
            if dp[i] > ans {
                ans = dp[i]
            }
        }
    }
    return ans
}