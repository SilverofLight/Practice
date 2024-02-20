package main

import "fmt"

func main() {
	var on [][]int
	for i := 0; i < 3; i++ {
		on = append(on, []int(nil))
		for j := 0; j < 3; j++ {
			on[i] = append(on[i], 0)
		}
	}
	fmt.Println(on)
}
