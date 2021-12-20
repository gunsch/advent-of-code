package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

var SIZE = 100
var MARGIN = 400
var START = (MARGIN - SIZE) / 2

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	enhancementString := scanner.Text()
	scanner.Scan()
	fmt.Println(enhancementString)

	pixels := makePixels(MARGIN)

	row := START
	for scanner.Scan() {
		input := scanner.Text()
		for idx, pixel := range input {
			col := START + idx
			if pixel == '.' {
				pixels[row][col] = false
			} else if pixel == '#' {
				pixels[row][col] = true
			} else {
				log.Fatalf("wtf!\n")
			}
		}
		row++
	}

	printPixels(pixels)

	iterations := 50
	for i := 0; i < iterations; i++ {
		pixels = doIteration(enhancementString, pixels)
	}
	printPixels(pixels)
	fmt.Printf("%d pixels lit\n", countLitPixels(pixels))
}

func countLitPixels(pixels [][]bool) int {
	count := 0
	distance := START / 2
	fmt.Printf("counting from distancae = %d\n", distance)
	for i := distance; i < MARGIN-distance; i++ {
		for j := distance; j < MARGIN-distance; j++ {
			if pixels[i][j] {
				count++
			}
		}
	}
	return count
}

func doIteration(enhancementString string, pixels [][]bool) [][]bool {
	newPixels := makePixels(MARGIN)

	for i := 1; i < MARGIN-1; i++ {
		for j := 1; j < MARGIN-1; j++ {
			newPixelVal := calculatePixel(enhancementString, pixels, i, j)
			// fmt.Printf("%d/%d=%v\n", i, j, newPixelVal)
			newPixels[i][j] = newPixelVal
		}
	}

	return newPixels
}

func calculatePixel(enhancementString string, pixels [][]bool, i, j int) bool {
	if i <= 0 || j <= 0 || i >= len(pixels)-1 || j >= len(pixels)-1 {
		log.Fatalf("nope")
	}

	bitString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s",
		mp(pixels[i-1][j-1]),
		mp(pixels[i-1][j]),
		mp(pixels[i-1][j+1]),
		mp(pixels[i][j-1]),
		mp(pixels[i][j]),
		mp(pixels[i][j+1]),
		mp(pixels[i+1][j-1]),
		mp(pixels[i+1][j]),
		mp(pixels[i+1][j+1]))
	number, _ := strconv.ParseInt(bitString, 2, len(bitString) + 1)
	// fmt.Printf("bitstring %s, number %d\n", bitString, number)

	return readBool(enhancementString, int(number))
}

func mp(v bool) string {
	if v {
		return "1"
	} else {
		return "0"
	}
}

func makePixels(size int) [][]bool {
	pixels := make([][]bool, size)
	for i := 0; i < size; i++ {
		pixels[i] = make([]bool, size)
	}
	return pixels
}

func readBool(enhancementString string, index int) bool {
	return []rune(enhancementString)[index] == '#'
}

func printPixels(pixels [][]bool) {
	for _, x := range pixels {
		for _, y := range x {
			char := '.'
			if y {
				char = '#'
			}
			fmt.Printf("%s", string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}
