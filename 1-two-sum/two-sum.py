class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        for i in range(len(nums)):
            tmp = target - nums[i]
            if tmp in nums:
                j = nums.index(tmp)
                if i != j:
                    return [i, j]