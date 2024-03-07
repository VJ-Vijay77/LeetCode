package main

import (
	"fmt"
	"sort"
)

func main() {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}
	// Output: [[1,6],[3,9],[4,5]]
	// Explanation:
	// The item with value = 1 occurs in items1 with weight = 1 and in items2 with weight = 5, total weight = 1 + 5 = 6.
	// The item with value = 3 occurs in items1 with weight = 8 and in items2 with weight = 1, total weight = 8 + 1 = 9.
	// The item with value = 4 occurs in items1 with weight = 5, total weight = 5.
	// Therefore, we return [[1,6],[3,9],[4,5]].
	fmt.Println("Result := ", mergeSimilarItems(items1, items2))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hm := make(map[int]int)
	var ans [][]int
	for _, item := range items1 {
	  hm[item[0]] += item[1]
	}
	for _, item := range items2 {
	  hm[item[0]] += item[1]
	}
  
	for k, v := range hm {
	  ans = append(ans, []int{k, v})
	}
  
	sort.Slice(ans, func(i, j int) bool {
	  return ans[i][0] < ans[j][0]
	})
  //smalll changes
	return ans
}
