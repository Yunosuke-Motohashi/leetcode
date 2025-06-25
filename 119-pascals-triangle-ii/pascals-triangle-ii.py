class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        dp = [[0]*i for i in range(1,rowIndex+2)]
        dp[0][0] = 1
        for i in range(1, rowIndex+1):
            for j in range(i+1):
                # print(i,j)
                if j == 0 or j == i:
                    dp[i][j] = 1
                else:
                    dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
        # print(dp)
        return dp[-1]
        