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
	// result数组是“最长相等前缀后缀数组”
	// result[i]表示着，以pattern[i]为结尾的后缀，最长的相同的前缀的长度
	// 比如对于pattern "abbcabb"，result数组为[-1, 0, 0, 0, 1, 2, 3]
	// 比如我们截取子串: 'abbcab', result[5], 表示了如果以下标5为结尾，那么目前子串前缀后缀最长相等的长度为2，即"ab"
	// 为啥非忒从减一开始呢？主要是出于使用上的考虑
	result[0] = front
	for i := 1; i < len; i += 1 {
		// 双指针，front指向前缀，i遍历pattern串，相当于遍历后缀结尾
		for front >= 0 && pattern[i] != pattern[front+1] {
			// 这里为啥是这样回退呢？而且为啥是i和front+1比较，不是i和front？
			// front目前由于是从-1开始的，指向的是并不是前缀的最后一个字符，而是倒数第二个，同时也是前缀的长度-1
			// 如果当前前缀的最后一个字符和当前后缀的最后一个字符不匹配，那我们就要回退前缀
			// 比如当前前缀为“aaab”，当前后缀为“aaac”，我们可以把前缀看作模式串，后缀看作待匹配串
			// ‘b’和‘c’不匹配的时候，我们要回退模式串的下标
			// 
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