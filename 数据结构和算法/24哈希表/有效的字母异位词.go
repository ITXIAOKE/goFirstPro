package _4哈希表

/**
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。


示例1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

*/

func isAnagram(s string, t string) bool {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}

	for _, k := range t {
		m[k-'a']--
	}

	for _, w := range m {
		if w != 0 {
			return false
		}
	}

	return true
}
