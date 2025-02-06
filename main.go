package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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

func main() {
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

	fmt.Println("Words in the book: ", num_words)
	fmt.Println("Character counts:")
	fmt.Println(count_per_char)
}
