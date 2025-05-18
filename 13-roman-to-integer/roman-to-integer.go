import "fmt"
import "strings"

func romanToInt(s string) int {

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
    if len(s) == 1 {
        return romanMap[s]
    }
    result := 0
    preVal := 0
    Val := 0

    for i := len(s)-1; i>=0; i-- {
        key := string(s[i])
        Val, _ = romanMap[string(key)]
        fmt.Println(result, preVal, Val)
        if Val >= preVal{
            result += Val
        } else {
            result -= Val
        }
        
        preVal = Val
    }

	fmt.Println(romanMap)
	return result
}