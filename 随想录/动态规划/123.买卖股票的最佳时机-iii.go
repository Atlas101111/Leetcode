/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (58.08%)
 * Likes:    1634
 * Dislikes: 0
 * Total Accepted:    302.7K
 * Total Submissions: 505K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：prices = [3,3,5,0,0,3,1,4]
 * 输出：6
 * 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 *
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
 *
 * 示例 4：
 *
 *
 * 输入：prices = [1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

// @lc code=start
func maxProfit(prices []int) int {
	k := 2
	dp := make([][]int, len(prices))
	for index, _ := range dp {
		dp[index] = make([]int, k*2)
	}

	for j := 0; j < k*2; j++ {
		if j%2 == 1 {
			dp[0][j] = -prices[0]
		}
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < k*2; j++ {
			if j == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1]+prices[i])
			} else if j == 1 {
				// 这里，如果写dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])就会过不了default样例
				dp[i][j] = max(dp[i-1][j], -prices[i])
			} else if j%2 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-3]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j+1]+prices[i])
			}
		}
	}

	// for i := 0; i < k*2; i++ {
	// 	fmt.Printf("%d\n", dp[7][i])
	// }

	return dp[len(prices)-1][k*2-2]

}

// @lc code=end

