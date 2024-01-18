package main
import (
	"fmt"
	"strings"
)
/*
题目：Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
 return the length of last word (last word means the last appearing word if we loop from left to right) in the string.

If the last word does not exist, return 0.

Note: A word is defined as a maximal substring consisting of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func length(s string)int {
	//先要去除前后的空格
	s1 := strings.TrimSpace(s)
	//切片
	s2 := strings.Split(s1, " ")
	return len(s2[len(s2)-1])
}
func main(){
	fmt.Println(length("Hello World  "))
}