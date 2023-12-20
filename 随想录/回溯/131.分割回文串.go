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
var (
	result [][]string
	path   []string
)

func partition(s string) [][]string {
	result = make([][]string, 0)
	path = make([]string, 0)

	dfs(s)
	return result
}

func dfs(s string) {
	if len(s) == 0 {
		temp := make([]string, len(path))
		copy(temp, path)
		result = append(result, temp)
		return
	}
	for i := 0; i < len(s); i++ {
		if isPalindrome(s[0:i+1]) == false {
			continue
		}
		path = append(path, s[0:i+1])
		nextS := s[i+1:]
		dfs(nextS)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end

