package iteration

func ContainsAny(s, chars string) bool {
    for _, c := range chars {
        for _, sc := range s {
            if c == sc {
                return true
            }
        }
    }

    return false
}
