
func removeElement(nums []int, val int) int {
	for {
		found := false
		for i := 0; i < len(nums); i++ {
			if nums[i] == val {
				nums = append(nums[:i], nums[i+1:]...)
				found = true
			}
		}
		if !found {
			break
		}
	}
	return len(nums)
}