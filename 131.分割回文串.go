/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode.cn/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (73.36%)
 * Likes:    1355
 * Dislikes: 0
 * Total Accepted:    257.2K
 * Total Submissions: 350.7K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 * 回文串 是正着读和反着读都一样的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func partition(s string) [][]string {
	result := make([][]string, 0)
	path := make([]string, 0)

	var dfs func(input string, startIndex int)
	dfs = func(input string, startIndex int) {

		inputLen := len(input)
		if startIndex == inputLen {
			temp := make([]string, len(path), len(path)*2)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		// for x := startIndex; x < inputLen; x += 1 {
		// 这个外层循环完全没有用，因为字串的开头必须是input[0]，循环开始下标反而会导致最后的结果拼不成输入字符串
		// 比如输入‘aab’，会出现[‘a’，‘b’]这样的结果
		// }

		for subLen := 1; startIndex+subLen <= inputLen; subLen += 1 {
			subString := input[startIndex : startIndex+subLen]

			if isPalindrome(subString) {
				path = append(path, subString)
				dfs(input, startIndex+subLen)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s, 0)
	return result
}

func isPalindrome(s string) bool {

	for x, y := 0, len(s)-1; x < y; x += 1 {
		if s[x] != s[y] {
			return false
		}
		y -= 1
	}
	return true
}

// @lc code=end

