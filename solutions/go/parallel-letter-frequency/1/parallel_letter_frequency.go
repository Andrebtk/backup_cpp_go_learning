package parallelletterfrequency

import "unicode"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
    txt := []rune(text)
	f := make(FreqMap)
	
    for _, char := range txt {

		if !unicode.IsLetter(char) {
            continue 
        }
        
        charN := unicode.ToLower(char)
        
        _, ok := f[charN]
        if ok {
            f[charN] += 1
        } else {
            f[charN] = 1
        }
    }

    return f
    
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	f := make(FreqMap)
	messages := make(chan FreqMap)

    for _, text := range texts {
        go func(t string) {
            messages <- Frequency(t)
        }(text)
    }


    for range len(texts) {
        result := <- messages

        // rune, int
        for key, valRes := range result {
            f[key] += valRes
        }

    }

    return f
    
}
