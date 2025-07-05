// import "fmt"
func divisorGame(n int) bool {
    count := 0
    for n > 1 {
        x := 1
        if n % x == 0 {
            n = n-x
            // fmt.Println(n, count)
        }
        count += 1
    }
    // fmt.Println(n, count)
    if count % 2 == 0 {
        return false
    } else {
        return true
    }
}
