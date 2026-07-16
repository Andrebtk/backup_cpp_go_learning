package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        if char == '❗' {
            return "recommendation"
        } else if char == '🔍' {
            return "search"
        } else if char == '☀' {
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	res := []rune(log)
    for index, char := range res {
        if char == oldRune {
            res[index] = newRune
        } 
    }

    return string(res)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for range log {
        limit--
        if limit < 0 {
            return false
        }
    }

    return true
}
