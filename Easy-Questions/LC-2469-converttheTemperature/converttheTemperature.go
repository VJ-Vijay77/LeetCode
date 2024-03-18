package main

import "fmt"

func main() {
	// Kelvin = Celsius + 273.15
	// Fahrenheit = Celsius * 1.80 + 32.00
	 
	
	celsius := 36.50
	// Output: [309.65000,97.70000]
	// Explanation: Temperature at 36.50 Celsius converted in Kelvin is 309.65 and converted in Fahrenheit is 97.70.
	fmt.Println("Result := ",convertTemperature(celsius))
}


func convertTemperature(celsius float64) []float64 {
    var res []float64
	res = append(res,celsius+273.15,celsius*1.80+32 )
	return res
}