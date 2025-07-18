package algo


import (
	"fmt"
)

// Q1 1. Write a function that takes a slice of integers and returns the count of even numbers in it.
func countEvenNumbers(numbers []int) int {
	count := 0
	for _, num := range numbers {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

// Q2 - . Given a slice of strings, return a map showing how many times each word appears.
func countWords(words []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

// Q3 - 3. Write a function that checks if two strings are anagrams of each other.
func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)
	for _, ch := range s1 {
		counts[ch]++
	}
	for _, ch := range s2 {
		counts[ch]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	// Q1
	numbers := []int{2, 4, 6, 8, 10}
	evenCount := countEvenNumbers(numbers)
	fmt.Println("Count of even numbers:", evenCount)

	// Q2
	words := []string{"Sudan", "India", "Cyprus", "India", "Nigeria", "Turkiye"}
	wordCount := countWords(words)
	fmt.Println("Word count:", wordCount)

	// Q3
	word1 := "listen"
	word2 := "silent"
	if areAnagrams(word1, word2) {
		fmt.Println("These words are anagrams.")
	} else {
		fmt.Println("These words are not anagrams.")
	}
}
