package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type fish struct {
	value int
	count int
}

func printO(octopi [][]int) {
	for _, y := range octopi {
		for _, x := range y {
			fmt.Printf(fmt.Sprintf("%x", x))
		}
		fmt.Println()
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	octopi := make([][]int, 10)
	for i := range octopi {
		octopi[i] = make([]int, 10)
	}

	row := 0
	for scanner.Scan() {
		input := scanner.Text()

		for i, v := range input {
			octopi[row][i], _ = strconv.Atoi(string(v))
		}
		row++
	}
	printO(octopi)

	steps := 0
	flashes := 0
	for steps < 500 {
		flashes += updateOctopi(octopi)

		steps++
		fmt.Printf("Step %d / %d flashes\n", steps, flashes)
		printO(octopi)
	}
}

func updateOctopi(octopi [][]int) int {
	// Increase all by one
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			octopi[i][j]++
		}
	}

	// Then, any octopus with an energy level greater than 9 flashes.
	// This increases the energy level of all adjacent octopuses by 1,
	// including octopuses that are diagonally adjacent.
	// If this causes an octopus to have an energy level greater than 9,
	// it also flashes. This process continues as long as new octopuses keep
	// having their energy level increased beyond 9. (An octopus can only 
	// flash at most once per step.)
	flashes := 0
	for {
		o_f := flashOctopi(octopi)
		if o_f > 0 {
			// fmt.Printf("flash\n")
			flashes += o_f
		} else {
			break
		}
	}
	fmt.Printf("%d flashes\n", flashes)

	if flashes == 100 {
		log.Fatalf("got all 100!")
	}

	// Finally, any octopus that flashed during this step has its energy level
	// set to 0, as it used all of its energy to flash.
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (octopi[i][j] == -1) {
				octopi[i][j] = 0
			}
		}
	}

	return flashes
}

func flashOctopi(octopi [][]int) int {
	flashes := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (octopi[i][j] > 9) {
				flashAll(octopi, i, j)
				flashes++
			}
		}
	}
	return flashes
}

func flashAll(octopi [][]int, i, j int) {
	flashOne(octopi, i - 1, j - 1)
	flashOne(octopi, i - 1, j)
	flashOne(octopi, i - 1, j + 1)
	flashOne(octopi, i, j - 1)
	flashOne(octopi, i, j + 1)
	flashOne(octopi, i + 1, j - 1)
	flashOne(octopi, i + 1, j)
	flashOne(octopi, i + 1, j + 1)

	// Set center to special value -1 so that it won't flash again.
	octopi[i][j] = -1
}

func flashOne(octopi [][]int, i, j int) {
	if i < 0 || j < 0 || i > 9 || j > 9 {
		return
	}

	// Already flashed this step
	if octopi[i][j] == -1 {
		return
	}

	octopi[i][j]++
}
