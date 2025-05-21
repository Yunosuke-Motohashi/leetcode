import (
    "strings" 
    "fmt"
)
func isValid(s string) bool {
    Stuck := []rune{}
    flg := true
    if len(s) <= 1{
        return false
    }
    if !strings.ContainsAny(s, ")}]") {
        return false
    }
    for _, char := range s {
        if char == ')' {
            // pop(first out)
            if len(Stuck)==0 {
                flg = false
                continue
            }
            tmpChar := Stuck[0]
            Stuck = Stuck[1:]
            if tmpChar != '('  {
                flg = false
            }
        } else if  char == '}' {
            if len(Stuck)==0 {
                flg = false
                continue
            }
            tmpChar := Stuck[0]
            Stuck = Stuck[1:]
            if tmpChar != '{' {
                flg = false
            }

        }  else if  char == ']' {
            if len(Stuck)==0 {
                flg = false
                continue
            }
            tmpChar := Stuck[0]
            Stuck = Stuck[1:]
            if tmpChar != '[' {
                flg = false
            }
        } else {
            // last in
            Stuck = append([]rune{char}, Stuck...)
        }
        if flg == false {
            break
        }
    }
    if len(Stuck) > 0 {
        return false
    }
    return flg
}