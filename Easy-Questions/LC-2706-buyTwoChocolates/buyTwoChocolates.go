package main

import "fmt"

func main() {
	prices := []int{98,54,6,34,66,63,52,39}
	money := 62
	// Output: 0
	// Explanation: Purchase the chocolates priced at 1 and 2 units respectively. You will have 3 - 3 = 0 units of money afterwards. Thus, we return 0.
	fmt.Println("Result := ", buyChoco(prices, money))

}

func buyChoco(prices []int, money int) int {
    min1, min2 := 101, 101
    for _, price := range prices {
        if price < min1 {
            min2, min1 = min1, price
        } else if price < min2 {
            min2 = price
        }
    }
    if min1+min2 <= money {
        return money - (min1 + min2)
    }
    return money
}
