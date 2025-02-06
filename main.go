package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func count_words(passage string) int {
	words := strings.Fields(passage)

	return len(words)
}

func count_characters(passage string) map[string]int {
	result := make(map[string]int)

	characters := strings.Split(passage, "")

	for _, char := range characters {
		lowercase_char := strings.ToLower(char)

		value := result[lowercase_char]
		if value > 0 {
			result[lowercase_char] += 1
		} else {
			result[lowercase_char] = 1
		}
	}

	return result
}

func report(dict map[string]int) {
	for key, value := range dict {
		if unicode.IsLetter(rune(key[0])) {
			fmt.Printf("The '%s' character was found %d times\n", key, value)
		}
	}
}

func main() {
	book_path := "books/frankenstein.txt"

	file, err := os.Open("books/frankenstein.txt")
	if err != nil {
		fmt.Print(err)
	}

	book, err := io.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	book_string := string(book)

	num_words := count_words(book_string)

	count_per_char := count_characters(book_string)

	fmt.Printf("--- Begin report of %s ---\n", book_path)
	fmt.Printf("%d words found in the document\n\n", num_words)
	report(count_per_char)
}
