
func romanToInt(s string) int {

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(s) == 1 {
		return romanMap[s[0]]
	}
	result := 0
	preVal := 0
	Val := 0

	for i := len(s) - 1; i >= 0; i-- {
		Val, _ = romanMap[s[i]]
		if Val >= preVal {
			result += Val
		} else {
			result -= Val
		}

		preVal = Val
	}

	return result
}