package main

import "fmt"

func main() {
	nums := []byte{'1', '2', '3', '4'}
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
