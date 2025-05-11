func numberOfSteps(num int) int {
    ans := 0
    for num > 0 {
        ans += 1
        if num % 2 == 0 {
            num = num /2
        } else {
            num -= 1
        }
    }
    return ans
}