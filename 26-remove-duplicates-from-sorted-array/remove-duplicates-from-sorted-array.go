import "slices"

func removeDuplicates(nums []int) int {
    uniqNums := slices.Compact(nums)
    return len(uniqNums)
}