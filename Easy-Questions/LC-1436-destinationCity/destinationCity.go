package main

import "fmt"

func main() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	// Output: "A"
	// Explanation: All possible trips are:
	// "D" -> "B" -> "C" -> "A".
	// "B" -> "C" -> "A".
	// "C" -> "A".
	// "A".
	// Clearly the destination city is "A".
	fmt.Println("Result := ", destCity(paths))
}

func destCity(paths [][]string) string {
	sources := make(map[string]struct{})

	for i := range paths {
		sources[paths[i][0]] = struct{}{}
	}

	for i := range paths {
		if _, ok := sources[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}

	return ""
}
