// import "fmt"

func getLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
    lastGroup := groups[0]
    result := []string{words[0]}
    for i:=1; i<n; i++ {
        if groups[i] != lastGroup {
            result = append(result, words[i])
            lastGroup = groups[i]
        }
    }
    
	return result
}