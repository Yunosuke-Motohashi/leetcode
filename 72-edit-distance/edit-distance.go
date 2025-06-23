import "fmt"
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    if n == m && word1 == word2 {
        return 0
    }
    
    // DPテーブル（スライス）を作成
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    // ベースケース：空の文字列への変換
    for i := 0; i<=m ; i++ {
        dp[i][0] = i
    }
    for j := 0; j<=n; j++ {
        dp[0][j] = j
    }

    for i:=1; i<=m; i++{
        for j:=1; j<=n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(
                    dp[i-1][j-1]+1, // 置換
                    dp[i][j-1]+1,// 挿入
                    dp[i-1][j]+1,// 削除
                )
            }
            // fmt.Println(i,j, dp)
        }
    }
    return dp[m][n]
}