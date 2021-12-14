package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// regex_fwd := regexp.MustCompile(`forward (?P<dist>\d+)`)

	const size int = 12

	counts := [size]int{}
	counted_rows := 0
	all_rows := [][]int{}
	for scanner.Scan() {
		input_str := scanner.Text()
		new_row := []int{}
		for i, v := range input_str {
			val := int(v - '0')
			counts[i] += val
			new_row = append(new_row, val)
		}
		all_rows = append(all_rows, new_row)
		counted_rows++
		// fmt.Println(scanner.Text())
	}
	fmt.Println(counts)

	fmt.Println(all_rows)

	gamma_results := [size]int{}
	epsilon_results := [size]int{}
	oxygen_mcv := [size]int{}
	co2_lcv := [size]int{}
	for i, v := range counts {
		fmt.Printf("%d %d %d\n", i, v, counted_rows)
		if v > counted_rows/2 {
			gamma_results[i] = 1
			epsilon_results[i] = 0
		} else {
			gamma_results[i] = 0
			epsilon_results[i] = 1
		}

		if v == counted_rows/2 {
			oxygen_mcv[i] = 1
			co2_lcv[i] = 0
		} else if v > counted_rows/2 {
			oxygen_mcv[i] = 1
			co2_lcv[i] = 0
		} else {
			oxygen_mcv[i] = 0
			co2_lcv[i] = 1
		}
	}

	fmt.Println("Oxygen time")
	remaining_rows := all_rows
	for i := 0; i < size; i++ {
		// First, find the most common number
		bit := find_oxygen_bit(remaining_rows, i)
		// fmt.Printf("found bit %d\n", bit)

		remaining_rows = filter_rows(remaining_rows, i, bit)
		if len(remaining_rows) == 1 {
			fmt.Println(remaining_rows)
		}
	}

	fmt.Println("CO2 time")
	remaining_rows_2 := all_rows
	for i := 0; i < size; i++ {
		// First, find the most common number
		bit := find_oxygen_bit(remaining_rows_2, i)
		// Flip for CO2
		if bit == 0 {
			bit = 1
		} else {
			bit = 0
		}
		// fmt.Printf("found bit %d\n", bit)

		remaining_rows_2 = filter_rows(remaining_rows_2, i, bit)
		if len(remaining_rows_2) == 1 {
			fmt.Println(remaining_rows_2)
		}
	}

	fmt.Println(gamma_results)
	fmt.Println(epsilon_results)
	fmt.Println(oxygen_mcv)
	fmt.Println(co2_lcv)

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}

func find_oxygen_bit(all_rows [][]int, index int) int {
	// fmt.Printf("searching for oxygen bit in index %d\n", index)

	total := 0
	for _, row := range all_rows {
		total += row[index]
	}

	// fmt.Printf("finding... total=%d, len=%d", total, len(all_rows))
	if total * 2 == len(all_rows) {
		return 1;
	} else if total * 2 > len(all_rows) {
		return 1;
	} else {
		return 0;
	}
}

func filter_rows(all_rows [][]int, index int, bit int) [][]int {
	fmt.Printf("Filtering rows for index %d == %d\n", index, bit)
	rows_to_return := [][]int{}
	for _, row := range(all_rows) {
		if row[index] == bit {
			rows_to_return = append(rows_to_return, row)
		}
	}
	return rows_to_return
}
