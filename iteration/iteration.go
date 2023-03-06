package iteration

// Repeat repeats a character n times
func Repeat(char string, times int) string {
    res := ""
    for i := 0; i < times; i++ {
        res += char
    }
    return res
}
