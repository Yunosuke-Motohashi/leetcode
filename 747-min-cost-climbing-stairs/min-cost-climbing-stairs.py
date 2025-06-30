class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n = len(cost)
        dp = [0] * (n + 1)
        dp[1] = cost[0]
        for i in range(2, n):
            dp[i] = cost[i-1] + min(dp[i - 1], dp[i - 2])
            # print(i, dp, cost[i-1], dp[i - 1], dp[i - 2])
        dp[n] = min(dp[n-1], dp[n-2]+cost[n-1])
        return dp[n]
