package main

import "fmt"


func letterfreq(str []string) map[rune]int {
	freq := make(map[rune]int)
	for _, str := range str {
		for _, ch := range str {
			if ch >= 'a' && ch <= 'z' {
				freq[ch]++
			}
		}
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		if _, exists := freq[ch]; !exists {
			freq[ch] = 0
		}
	}

	return freq
}

func main() {
	str := []string{"quick", "brown", "fox", "lazy", "dog"}

	freq := letterfreq(str)

	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c: %d\n", ch, freq[ch])
	}
}
