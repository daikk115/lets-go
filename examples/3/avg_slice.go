/* Write a function that calculates the average of a float64 slice */

package main

import "fmt"

func avg(input_list []float64) (output float64) {
	var sum float64
	var len_input_list float64
	sum = 0
	len_input_list = float64(len(input_list))
	if len(input_list) != 0 {
		for _, v := range input_list {
			sum += v
		}
		output = sum / len_input_list
	}
	return
}

func main() {
	input_list := []float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println("Average: ", avg(input_list))
}
