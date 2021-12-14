package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type fish struct {
	value int
	count int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	total_score := 0
	scores := []int{}
	for scanner.Scan() {
		input := scanner.Text()

		current_score := match(input)
		total_score += current_score

		// break
		if current_score > 0 {
			scores = append(scores, current_score)
		}
	}	
	fmt.Println(total_score)
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Printf("middle: %d\n", scores[len(scores) / 2])
}

func match(s string) int {
	fmt.Printf("\n%s\n", s)

	// [({(<(())[]>[[{[]{<()<>>
	char_stack := []rune{}
	char_map := map[rune]rune {
		'(' : ')',
		'[' : ']',
		'{' : '}',
		'<' : '>',
	}
	score_map := map[rune]int {
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	// inverse_char_map := map[rune]rune {
	// 	')' : '(',
	// 	']' : '[',
	// 	'}' : '{',
	// 	'>' : '<',
	// }

	for _, char := range s {
		// Is it a starter? always okay.
		if ender, ok := char_map[char]; ok {
			char_stack = append(char_stack, ender)
			// fmt.Printf("Add, char stack: %s\n", string(char_stack))
			continue
		}

		// OK, it's an ender. Now what?
		// starter := inverse_char_map[char]
		expected := char_stack[len(char_stack) - 1]
		if expected == char {
			// Happy case! Remove the end.
			char_stack = char_stack[:len(char_stack) - 1]
		} else {
			fmt.Printf("fail! expected %s, found %s\n", string(expected), string(char))
			// Part 1
			// return score_map[char]

			// Part 2
			return score_map[char] * 0
		}

		// fmt.Printf("Rem, char stack: %s\n", string(char_stack))
	}
	fmt.Printf("Char stack: %s\n", string(char_stack))

	// Part 1
	// return 0
	// Part 2
	return score(char_stack)
}

func score(char_stack []rune) int {
	score_map := map[rune]int {
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	// Reverse
	for i, j := 0, len(char_stack)-1; i < j; i, j = i+1, j-1 {
		char_stack[i], char_stack[j] = char_stack[j], char_stack[i]
	}

	score := 0
	for _, v := range char_stack {
		score = score * 5
		score += score_map[v]
	}
	
	fmt.Printf("%s = %d\n", string(char_stack), score)
	return score
}
