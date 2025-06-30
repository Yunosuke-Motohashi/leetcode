class Solution:
    def fib(self, n: int) -> int:
        fibonacci_nums = [0, 1]
        if n >= 2:
            for i in range(2, n+1):
                fibonacci_nums.append(fibonacci_nums[i-2]+fibonacci_nums[i-1])
        return fibonacci_nums[n]


        