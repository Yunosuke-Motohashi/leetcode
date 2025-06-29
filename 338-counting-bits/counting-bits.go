import "fmt"

func countBits(n int) []int {
    ans := []int{}
    for i:= 0; i<=n; i++ {
        
        countOne := 0
        j := i
        for j > 0 {
            countOne += j & 1 // 最下位ビットをチェック（0 or 1）
            j >>= 1           // 右シフトして次のビットを処理
        }
        fmt.Println(countOne)
        ans = append(ans, countOne)
    }
    return ans
}
