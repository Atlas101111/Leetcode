/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (37.95%)
 * Likes:    1397
 * Dislikes: 0
 * Total Accepted:    378.9K
 * Total Submissions: 998.5K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	len := len(nums)
	result := make([][]int, 0)
	for left := 0; left < len; left += 1 {
		if nums[left] > target && nums[left] > 0 {
			break
		}

		if left > 0 && nums[left] == nums[left-1] {
			continue
		}

		for left2 := left + 1; left2 < len; left2 += 1 {
			right := left2 + 1
			right2 := len - 1

			if left2 > left+1 && nums[left2] == nums[left2-1] {
				continue
			}

			for right < right2 {
				temp := nums[left] + nums[left2] + nums[right] + nums[right2]

				if temp == target {
					x := []int{nums[left], nums[left2], nums[right], nums[right2]}
					result = append(result, x)

					for right < right2 && nums[right] == nums[right+1] {
						right += 1
					}

					for right < right2 && nums[right2] == nums[right2-1] {
						right2 -= 1
					}

					right += 1
					right2 -= 1
				} else if temp < target {
					right += 1
				} else {
					right2 -= 1
				}
			}

		}
	}
	return result
}

// @lc code=end

