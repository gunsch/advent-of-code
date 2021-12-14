package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type fish struct {
	value int
	count int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	directions := make(map[string][]string)
	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		path_start, path_end := path[0], path[1]
		directions[path_start] = append(directions[path_start], path_end)
		if (path_start != "start" && path_end != "end") {
			directions[path_end] = append(directions[path_end], path_start)
		}
	}
	fmt.Println(directions)
	
	// paths_to_work := []string{",start,"}
	completed_paths := map[string]bool{}
	paths_to_work := make([][]string, 1)
	current_path := []string{}
	paths_to_work[0] = []string{"start"}
	for len(paths_to_work) > 0 {
		current_path, paths_to_work = paths_to_work[0], paths_to_work[1:]
		current_cave := current_path[len(current_path)-1]
		fmt.Printf("current cave: %s\n", current_cave)
		fmt.Printf("current path: %s\n", current_path)

		// Completion step
		if current_cave == "end" {
			path_str := strings.Join(current_path, ",")
			if !completed_paths[path_str] {
				fmt.Printf("completed path: %s\n", path_str)
			}
			completed_paths[path_str] = true
			continue
		}

		for _, next_cave := range directions[current_cave] {
			// Force creation of a different path
			next_path := append(make([]string, 0), current_path...)

			// Can we actually go to the next cave?
			first_rune, _ := utf8.DecodeRuneInString(next_cave)
			if unicode.IsLower(first_rune) && contains(current_path, next_cave) {
				// Thennn this path isn't going to work.
				// fmt.Printf("Skipping %s, already visited in %s\n", next_cave, current_path)

				// Okay, ONE chance. Special case.
				if !contains(current_path, "DOUBLE") {
					next_path = append(next_path, "DOUBLE")
				} else {
					continue
				}
			}

			next_path = append(next_path, next_cave)
			// fmt.Printf("appending %s to current path = %s, pointer %v\n", next_cave, next_path, reflect.ValueOf(next_path).Pointer())
			paths_to_work = append(paths_to_work, next_path)
		}

		fmt.Printf("paths remaining to work: %d\n", len(paths_to_work))
	}

	fmt.Printf("completed paths: %d\n", len(completed_paths))
	
	paths_arr := []string{}
	for path, _ := range completed_paths {
		// fmt.Println(path)
		paths_arr = append(paths_arr, path)
	}
	sort.Strings(paths_arr)
	for _, path := range paths_arr {
		fmt.Println(path)
	}
	fmt.Printf("completed paths: %d\n", len(completed_paths))


}

func contains_twice(s[]string, e string) bool {
	count := 0
    for _, a := range s {
        if a == e {
            count++
			if count == 2 {
				return true
			}
		}
    }
    return false
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

