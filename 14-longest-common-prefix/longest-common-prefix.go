
func longestCommonPrefix(strs []string) string {
    shortestStrCount := 201
    for _, str := range strs {
        shortestStrCount = min(shortestStrCount, len(str))
    }
    ans := ""
    for i:=0;i<shortestStrCount;i++ {
        preChar := strs[0][i]
        flg := true
        for _, str := range(strs) {
            currChar := str[i]
            if preChar != currChar {
                flg = false
                break
            }
        }
        if flg == true {
            ans += string(preChar)
        } else {
            break
        }
    }
    return ans

}