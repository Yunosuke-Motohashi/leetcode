func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, banks := range accounts {
        tmpWealth := 0
        for _, wealth := range banks {
            tmpWealth += wealth
        }
        if maxWealth < tmpWealth {
            maxWealth = tmpWealth
        }
    }
    return maxWealth
}