package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	// Output: 4
	// Explanation: The students are moved as follows:
	// - The first student is moved from from position 2 to position 1 using 1 move.
	// - The second student is moved from from position 7 to position 5 using 2 moves.
	// - The third student is moved from from position 4 to position 3 using 1 move.
	// In total, 1 + 2 + 1 = 4 moves were used.
	fmt.Println("Result := ", minMovesToSeat(seats, students))
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var count float64
	for i := 0; i < len(seats); i++ {
		count += math.Abs(float64(seats[i]-students[i]))
	}

	return int(count)
}
