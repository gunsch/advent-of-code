package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type rule struct {
	l1 string
	l2 string
	r1 string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	regex_insertion := regexp.MustCompile(`(?P<x1>[A-Z][A-Z]) -> (?P<x3>[A-Z])`)

	// Template line
	scanner.Scan()
	template := scanner.Text()
	// Blank line
	scanner.Scan()

	rules := make(map[string]rune)
	for scanner.Scan() {
		input := scanner.Text()

		if regex_insertion.MatchString(input) {
			matches := regex_insertion.FindStringSubmatch(input)
			rules[matches[1]] = []rune(matches[2])[0]
			continue
		}
	}

	fmt.Println(template)
	fmt.Println(rules)

	template_copy := template
	steps := 0
	for steps < 10 {
		template_copy = doInsertion([]rune(template_copy), rules)
		// fmt.Println(template_copy)
		fmt.Printf("%d / size: %d\n", steps, len(template_copy))
		steps++
		counts := countChars(template_copy)
		fmt.Println(counts)
	}

	start_map := make(map[string]int)
	for i := 0; i < len(template) - 1; i++ {
		start_map[string(template[i:i+2])]++
	}
	fmt.Println("fast path")
	fmt.Println(start_map)

	steps2 := 0
	for steps2 < 40 {
		start_map = doInsertionFast(start_map, rules)
		fmt.Println(start_map)
		steps2++
		counts := countChars2(template, start_map)
		fmt.Printf("round %d\n", steps2)
		fmt.Println(counts)
	}

}

func countChars2(template string, start_map map[string]int) []int {
	counts := make(map[rune]int)

	// This double-counts everything except first and last
	for pair, count := range start_map {
		pair_r := []rune(pair)
		counts[pair_r[0]] += count
		counts[pair_r[1]] += count
	}

	// Re-add first and last before dividing in half
	template_r := []rune(template)
	counts[template_r[0]]++
	counts[template_r[len(template_r) - 1]]++

	counts_i := []int{}
	for i, v := range counts {
		counts[i] = v / 2
		counts_i = append(counts_i, v / 2)
	}

	sort.Ints(counts_i)
	return counts_i
}

// NNCB --> 
// 1xNN 1xNC 1xCB
// --> C, B, H
// NC CN, NB BC, CH HB

func doInsertionFast(start map[string]int, rules map[string]rune) map[string]int {
	new_scores := make(map[string]int)

	for pair, value := range start {
		str1 := string([]rune{[]rune(pair)[0], rules[pair]})
		str2 := string([]rune{rules[pair], []rune(pair)[1]})
		new_scores[str1] += value
		new_scores[str2] += value
	}

	return new_scores
}

func doInsertion(template []rune, rules map[string]rune) string {
	extras := []string {}
	new_string := make([]rune, len(template) * 2 - 1)
	for i := 0; i < len(template) - 1; i++ {
		extras = append(extras, )
		new_string[i * 2] = template[i]
		new_string[i * 2 + 1] = rules[string(template[i:i+2])]
	}
	new_string[2*(len(template)-1)] = template[len(template)-1]
	return string(new_string)
}

func countChars(template string) map[rune]int {
	counts := make(map[rune]int)
	for _, v := range template {
		counts[v]++
	}
	return counts
}
