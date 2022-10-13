/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Medium (41.61%)
 * Likes:    1617
 * Dislikes: 0
 * Total Accepted:    747.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}
	lps := getLPS(needle)

	j := -1
	for i := 0; i < haystackLen; i += 1 {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = lps[j]
		}
		if haystack[i] == needle[j+1] {
			j += 1
		}
		if j == needleLen-1 {
			return i - needleLen + 1
		}
	}

	return -1
}

func getLPS(pattern string) []int {
	len := len(pattern)
	result := make([]int, 10005, 10010)
	front := -1
	result[0] = front
	for i := 1; i < len; i += 1 {
		for front >= 0 && pattern[i] != pattern[front+1] {
			front = result[front]
		}

		if pattern[i] == pattern[front+1] {
			front += 1
		}
		
		result[i] = front
	}
	return result
}

// @lc code=end

`