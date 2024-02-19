package main

import "fmt"

/*给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:

输入: strs = [""]
输出: [[""]]

示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/

//通过列哈希表，把每个字母表示成一个质数求乘积
func groupAnagrams(strs []string) [][]string {
	prime := map[byte]int{
		'a': 2,
		'b': 3,
		'c': 5,
		'd': 7,
		'e': 11,
		'f': 13,
		'g': 17,
		'h': 19,
		'i': 23,
		'j': 29,
		'k': 31,
		'l': 37,
		'm': 41,
		'n': 43,
		'o': 47,
		'p': 53,
		'q': 59,
		'r': 61,
		's': 67,
		't': 71,
		'u': 73,
		'v': 79,
		'w': 83,
		'x': 89,
		'y': 97,
		'z': 101,
	}
	path := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		s := 1
		for j := 0; j < len(strs[i]); j++ {
			s = s * prime[strs[i][j]]
		}
		path[i] = s
	}
	var ans [][]string
	m := make(map[int]bool)
	for i := 0; i < len(strs); i++ {

		ans = append(ans, []string(nil))

		for j := i; j < len(strs); j++ {
			if path[i] == path[j] && !m[path[i]] {
				ans[i] = append(ans[i], strs[j])
			}
		}
		m[path[i]] = true
	}
	ans2 := [][]string{}
	for i := 0; i < len(strs); i++ {
		if ans[i] != nil {
			ans2 = append(ans2, ans[i])
		}
	}
	return ans2
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
