/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (54.49%)
 * Likes:    3420
 * Dislikes: 0
 * Total Accepted:    804.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */

// @lc code=start

// DP做法
func lengthOfLIS_DP(nums []int) int {
	numsLen := len(nums)
	dp := make([]int, numsLen)

	for index, num := range nums {
		dp[index] = 1
		if index == 0 {
			continue
		}

		for current := index - 1; current >= 0; current-- {
			if nums[current] < num && dp[current]+1 > dp[index] {
				dp[index] = dp[current] + 1
			}
		}
	}

	max := 1
	for _, num := range dp {
		if num >= max {
			max = num
		}
	}

	return max
}

// 贪心做法
func lengthOfLIS(nums []int) int {
	result := make([]int, 0)

	for index, num := range nums {
		resultLen := len(result)
		if index == 0 || num > result[resultLen-1] {
			result = append(result, num)
			continue
		} else {
			first := binaryFindMin(result, num)
			result[first] = num
		}
	}
	return len(result)
}

// 返回大于等于target的第一个值的下标
func binaryFindMin(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// @lc code=end

