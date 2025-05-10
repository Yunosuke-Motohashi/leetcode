func runningSum(nums []int) []int {
    runningSum := []int{}
    sum := 0
    for _, i := range nums {
        sum += i
        runningSum = append(runningSum, sum)
    }
    return runningSum
}