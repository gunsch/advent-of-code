package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fwd_position := 0
	depth_position := 0
	aim := 0

	scanner := bufio.NewScanner(os.Stdin)
	regex_fwd := regexp.MustCompile(`forward (?P<dist>\d+)`)
	regex_up := regexp.MustCompile(`up (?P<dist>\d+)`)
	regex_down := regexp.MustCompile(`down (?P<dist>\d+)`)
	for scanner.Scan() {
		direction_cmd := scanner.Text()
		if z := regex_fwd.FindString(direction_cmd); z != "" {
			distance_str := regex_fwd.FindStringSubmatch(direction_cmd)[1]
			distance, _ := strconv.Atoi(distance_str)
			fwd_position += distance
			depth_position += (aim * distance)
		} else if z := regex_up.FindString(direction_cmd); z != "" {
			distance_str := regex_up.FindStringSubmatch(direction_cmd)[1]
			distance, _ := strconv.Atoi(distance_str)
			// depth_position -= distance
			aim -= distance
			if depth_position < 0 {
				log.Fatalf("Depth <0: %d\n", depth_position)
			}
		} else if z := regex_down.FindString(direction_cmd); z != "" {
			distance_str := regex_down.FindStringSubmatch(direction_cmd)[1]
			distance, _ := strconv.Atoi(distance_str)
			// depth_position += distance
			aim += distance
		} else {
			log.Fatalf("unhandled input: %s\n", scanner.Text())
		}

		fmt.Println(scanner.Text())
		fmt.Printf("Status: horizontal=%d, depth=%d, aim=%d\n", fwd_position, depth_position, aim)

	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	fmt.Printf("Forward position: %d\n", fwd_position)
	fmt.Printf("Depth: %d\n", depth_position)
	fmt.Printf("Total: %d\n", depth_position*fwd_position)
}
