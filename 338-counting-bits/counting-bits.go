// import "fmt"

func countBits(n int) []int {
    ans := make([]int, n+1)
    for i:= 1; i<=n; i++ {
        ans[i] += ans[i>>1]
        ans[i] += i & 1
        // fmt.Println(i, i>>1, ans)
    }
    return ans
}
