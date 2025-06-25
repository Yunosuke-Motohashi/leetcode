import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	} else {
		PreRow := []int{1}
        NextRow := []int{}

        for i:=1; i<=rowIndex; i++ {
		    NextRow = make([]int, i+1)
            for j:=0; j<len(NextRow);j++ {
                fmt.Println(i, j, PreRow, NextRow)
                if j==0 || j == i {
                    NextRow[j] = 1
                } else {
                    NextRow[j] = PreRow[j-1]+PreRow[j]
                }
            }
            PreRow = NextRow
		}
        return NextRow
	}
}