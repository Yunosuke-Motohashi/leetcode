import (
    "strconv"
)
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    xs := strconv.Itoa(x)
    l := 0
    r := len(xs) -1
    flg := true
    for l<r {
        if xs[l] != xs[r] {
            flg = false
            break
        }
        l ++
        r --
    }
    return flg
}