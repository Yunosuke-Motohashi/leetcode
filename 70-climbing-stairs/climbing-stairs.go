// import "fmt"
func climbStairs(n int) int {
    ans := []int{1, 1}
    if n <= 1 {
        return ans[n]
    }

    for i := 2; i < n+1; i++ {
        tmp := ans[i-1] + ans[i-2]
        ans = append(ans, tmp)
        // fmt.Println(i, ans)

    }
    return ans[n]
}