package _4哈希表

/**
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false

示例 2：
输入：ransomNote = "aa", magazine = "ab"
输出：false

示例 3：
输入：ransomNote = "aa", magazine = "aab"
输出：true


*/
func canConstruct(ransomNote string, magazine string) bool {
	var mag [26]int

	for _, v := range magazine {
		mag[v-'a']++
	}

	for _, k := range ransomNote {
		if mag[k-'a'] <= 0 {
			return false
		} else {
			mag[k-'a']--
		}
	}
	return true
}
