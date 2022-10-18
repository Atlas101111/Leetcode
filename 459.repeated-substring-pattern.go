/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode.cn/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (51.09%)
 * Likes:    804
 * Dislikes: 0
 * Total Accepted:    146.9K
 * Total Submissions: 287.7K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abab"
 * 输出: true
 * 解释: 可由子串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aba"
 * 输出: false
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "abcabcabcabc"
 * 输出: true
 * 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 10^4
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	inputLen := len(s)
	if inputLen == 1 {
		return false
	}

	lps := getNext(s)

	last := lps[inputLen-1]
	perfixLen := last + 1
	remainLen := inputLen - perfixLen

	if perfixLen >= remainLen && perfixLen%remainLen == 0 {
		return true
	}

	return false
}

func getNext(pattern string) []int {
	result := make([]int, 10010)
	j := -1
	result[0] = j

	for i := 1; i < len(pattern); i += 1 {
		for j >= 0 && pattern[i] != pattern[j+1] {
			j = result[j]
		}

		if pattern[i] == pattern[j+1] {
			j += 1
		}

		result[i] = j
	}

	return result

}

// @lc code=end

