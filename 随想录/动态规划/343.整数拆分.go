/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode.cn/problems/integer-break/description/
 *
 * algorithms
 * Medium (62.35%)
 * Likes:    1300
 * Dislikes: 0
 * Total Accepted:    290K
 * Total Submissions: 463.5K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
 *
 * 返回 你可以获得的最大乘积 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 *
 * 输入: n = 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= n <= 58
 *
 *
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}

	return dp[n]
}

// @lc code=end

