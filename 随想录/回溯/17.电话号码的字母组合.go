/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (57.99%)
 * Likes:    2268
 * Dislikes: 0
 * Total Accepted:    622K
 * Total Submissions: 1.1M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */

// @lc code=start

var (
	dict   map[string]string
	queue  []string
	result []string
	path   string
)

func letterCombinations(digits string) []string {
	result, queue = make([]string, 0), make([]string, 0)

	dict = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	for i := 0; i < len(digits); i++ {
		queue = append(queue, dict[digits[i:i+1]])
	}
	if len(queue) > 0 {
		dfs(queue)
	}
	return result
}

func dfs(remain []string) {
	if len(remain) == 0 {
		temp := make([]byte, len(path))
		copy(temp, path)
		result = append(result, string(temp))
		return
	}

	target := remain[0]
	for i := 0; i < len(target); i++ {
		path += target[i : i+1]
		remain = remain[1:]
		dfs(remain)
		remain = append([]string{target}, remain...)
		path = path[:len(path)-1]
	}
}

// @lc code=end

