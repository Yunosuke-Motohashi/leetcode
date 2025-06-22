class Solution:
    def climbStairs(self, n: int) -> int:
        ans = [1, 1]
        if n>=0:
            for i in range(2, n+1):
                ans.append(ans[i-2]+ans[i-1])
        return ans[n]
        