class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m, n = len(word1), len(word2)
        if word1 == word2:
            return 0
        dp = [[0]*(n+1) for _ in range(m+1)]
        # base case
        for i in range(m + 1):
            dp[i][0] = i  # 削除のみ
        for j in range(n + 1):
            dp[0][j] = j  # 挿入のみ


        for i in range(1,m+1):
            for j in range(1, n+1):
                if word1[i-1] == word2[j-1]:
                    dp[i][j] = dp[i-1][j-1]
                else:
                    replace = dp[i-1][j-1] + 1
                    insert = dp[i][j-1] + 1
                    delete = dp[i-1][j] + 1
                    dp[i][j] = min(replace, insert, delete)
        return dp[m][n]