package stringutil

func reverseTwo(s string) string {
    
    str := []rune(s)
    strLen := len(str)
    for i, j := 0, strLen-1; i < strLen/2; i , j = i+1, j -1 {
        str[i], str[j] = str[j], str[i]
    }
    return string(str)
}