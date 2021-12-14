package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func nonzero(a int) int {
	if a < 0 {
		return 0
	}
	return a
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	last_input := 100000000
	number_of_increases := 0

	input_sequence := []int{}
	last_sum := 10000000
	number_of_sum_increases := 0

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)

		input_num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Expected int, got: %s\n", input)
		}
		input_sequence = append(input_sequence, input_num)

		if len(input_sequence) >= 3 {
			last_three := input_sequence[nonzero(len(input_sequence)-3):]
			sum_of_three := sum(last_three)
			fmt.Println(last_three)
			fmt.Printf("sum: %d\n", sum_of_three)
			if sum_of_three > last_sum {
				number_of_sum_increases++
			}
			last_sum = sum_of_three
		}

		if input_num > last_input {
			number_of_increases++
		}
		last_input = input_num
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("Number of increases: %d\n", number_of_increases)
	fmt.Printf("Number of sum increases: %d\n", number_of_sum_increases)

}
