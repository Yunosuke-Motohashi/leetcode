// import "fmt"
func divisorGame(n int) bool {
    dp := make([]bool, 1001)
    dp[0], dp[1] = false, false
    for i :=2; i<n+1; i++ {
        for j:=1; j*j < i; j++ {
            if i%j == 0 {
                if !dp[i-j] {
                    dp[i] = true
                    break
                }
                if !dp[i-(i/j)] && j != i/j && i != i/j {
                    dp[i] = true
                    break
                }
            }
        }
        // fmt.Println(dp[1:n+1])
    }
    return dp[n]
}
// func divisorGame(n int) bool {
//     count := 0
//     for n > 1 {
//         x := 1
//         if n % x == 0 {
//             n = n-x
//             // fmt.Println(n, count)
//         }
//         count += 1
//     }
//     // fmt.Println(n, count)
//     if count % 2 == 0 {
//         return false
//     } else {
//         return true
//     }
// }
