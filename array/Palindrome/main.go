package main
import "fmt"
import "math"
/*
题目：Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
func getPalin(s string, l int, r int)int {//输入字符串、中心位置，返回最大的回文长度
    for l >=0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    return r - l - 1
}
func longest(s string)string {
    if len(s) < 1 {
        return ""
    }
    var end int
    var start int//用来记录最大的回文长度和位置

    for i := 0; i < len(s); i++ {
        //以i为中心散开
        s1 := getPalin(s, i, i)
        s2 := getPalin(s, i, i+1)//回文可能是aba或abba型
        tmp := int(math.Max(float64(s1),float64(s2)))
        if tmp > end - start {
            start = i - (tmp-1)/2
            end = i + tmp/2
        }
    }
    return s[start : end + 1]

}
func main(){
    fmt.Println(longest("babad"))
}