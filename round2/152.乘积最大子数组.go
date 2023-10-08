/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode.cn/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (43.10%)
 * Likes:    2116
 * Dislikes: 0
 * Total Accepted:    383.4K
 * Total Submissions: 889.3K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 *
 * 测试用例的答案是一个 32-位 整数。
 *
 * 子数组 是数组的连续子序列。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
 *
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	numsLen := len(nums)
	dpMax := make([]int, numsLen)
	dpMin := make([]int, numsLen)
	max := nums[0]

	for index, num := range nums {
		dpMax[index] = num
		dpMin[index] = num
		if index == 0 {
			continue
		}

		tempMax := num
		tempMin := num
		if num > 0 {
			tempMax = num * dpMax[index-1]
			tempMin = num * dpMin[index-1]
		} else {
			tempMax = num * dpMin[index-1]
			tempMin = num * dpMax[index-1]
		}

		if tempMax > num {
			dpMax[index] = tempMax
		}
		if tempMin < num {
			dpMin[index] = tempMin
		}

		if dpMax[index] > max {
			max = dpMax[index]
		}
	}

	return max
}

// @lc code=end

