package main
import "fmt"

/*字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

输入："aabcccccaaa"
输出："a2b1c5a3"
*/

func compressString(s string)string {
	if s == ""{
		return ""
	}
	res := ""
	count := 0
	for i := 1; i < len(s); i++ {
		count++
		if s[i-1] != s[i] {
			res += string(s[i-1]) + fmt.Sprintf("%d", count)
			count = 0
		}
	}
	res += string(s[len(s)-1]) + fmt.Sprintf("%d", count+1)
	if len(res) >= len(s){
		return s
	}else {
		return res
	}
}

func main(){
	fmt.Println(compressString("abc"))
}