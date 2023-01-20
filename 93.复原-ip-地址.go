/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode.cn/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (57.69%)
 * Likes:    1119
 * Dislikes: 0
 * Total Accepted:    298.9K
 * Total Submissions: 518K
 * Testcase Example:  '"25525511135"'
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 * 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能
 * 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅由数字组成
 *
 *
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	path := make([]byte, 0)

	var dfs func(input string, startIndex int, segmentCount int)
	dfs = func(input string, startIndex int, segmentCount int) {
		inputLen := len(input)
		if startIndex == inputLen && segmentCount == 4 {
			temp := make([]byte, len(path), len(path)*2)
			copy(temp, path)
			// Remove trailing '.'
			temp = temp[:len(temp)-1]
			result = append(result, string(temp))
			return
		}

		for subLen := 1; subLen+startIndex <= inputLen; subLen += 1 {
			subString := input[startIndex : subLen+startIndex]
			if isValid(subString) {
				path = append(path, []byte(subString)...)
				path = append(path, '.')
				dfs(input, startIndex+subLen, segmentCount+1)
				path = path[:len(path)-len(subString)-1]
			}
		}
	}
	dfs(s, 0, 0)
	return result
}

func isValid(input string) bool {
	if len(input) > 1 && input[0] == '0' {
		return false
	}

	digit, _ := strconv.Atoi(input)
	if digit < 0 || digit > 255 {
		return false
	}

	return true
}

// @lc code=end

