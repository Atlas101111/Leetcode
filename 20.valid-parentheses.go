/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.53%)
 * Likes:    3559
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.9M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 每个右括号都有一个对应的相同类型的左括号。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 0)
	result := true
	for _, char := range s {
		if isLeftBracket(char) {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if isMatch(top, char) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return result
}

func isLeftBracket(char rune) bool {
	if char == '[' || char == '{' || char == '(' {
		return true
	}
	return false
}

func isMatch(left rune, right rune) bool {
	if left == '[' && right == ']' {
		return true
	} else if left == '(' && right == ')' {
		return true
	} else if left == '{' && right == '}' {
		return true
	}
	return false
}

// @lc code=end

