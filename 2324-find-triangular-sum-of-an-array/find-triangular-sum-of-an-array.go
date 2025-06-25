// import "fmt"
func triangularSum(nums []int) int {
    next := []int{}
    for len(nums) != 1 {
        next = []int{}
        for i := 0; i<len(nums)-1; i++ {
            next = append(next, (nums[i]+nums[i+1])%10)
        }
        // fmt.Println(next)
        nums = next
    }
    return nums[0]
}