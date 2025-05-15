
func twoSum(nums []int, target int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			tmp := nums[i] + nums[j]
			if i != j && tmp == target {
				ans = []int{i, j}
				break
			}
		}
		if len(ans) > 0 {
			break
		}
	}
	return ans
}