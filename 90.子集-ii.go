/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode.cn/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.67%)
 * Likes:    1003
 * Dislikes: 0
 * Total Accepted:    267K
 * Total Submissions: 419.4K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 *
 *
 *
 *
 */

// @lc code=start
// func subsetsWithDup(nums []int) [][]int {
// 	result := make([][]int, 0)
// 	path := make([]int, 0)
// 	result = append(result, path)
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	var dfs func(input []int, used []bool, endIndex int)
// 	dfs = func(input []int, used []bool, endIndex int) {
// 		if endIndex >= len(input) {
// 			return
// 		}

// 		if endIndex > 0 && input[endIndex] == input[endIndex-1] && used[endIndex-1] {
// 			used[endIndex] = true
// 			dfs(input, used, endIndex+1)
// 			return
// 		}

// 		preResult := make([][]int, len(result), 2*len(result))
// 		copy(preResult, result)
// 		for _, item := range preResult {
// 			// if endIndex >= len(input) {
// 			// 	return
// 			// }

// 			temp := make([]int, len(item), 2*len(item))
// 			copy(temp, item)

// 			temp = append(temp, input[endIndex])
// 			result = append(result, temp)

// 			current := endIndex
// 			for step := 1; current+step <= len(input)-1 && input[current] == input[current+step]; step += 1 {
// 				token := make([]int, len(temp), 2*len(temp))
// 				copy(token, temp)

// 				token = append(token, input[current+1:current+step+1]...)
// 				result = append(result, token)
// 				// endIndex = current + step
// 			}

// 		}

// 		used[endIndex] = true
// 		dfs(input, used, endIndex+1)
// 	}

// 	dfs(nums, make([]bool, len(nums), 2*len(nums)), 0)
// 	return result
// }

// 下面是数层遍历的做法
func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	result = append(result, path)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var dfs func(input []int, used []bool, startIndex int)
	dfs = func(input []int, used []bool, startIndex int) {
		if startIndex >= len(input) {
			return
		}

		for i := startIndex; i < len(input); i += 1 {
			// 这里要判断used[i-1]是false
			// 和leetcode40 组合排序是一样的
			// 因为要保证的是同一深度中的去重，如果used[i-1]是true，说明是上一层传下来的
			// 同一层中的used数组，从startIndex到input末尾，只有一个是true的
			if i > startIndex && input[i] == input[i-1] && used[i-1] == false {
				continue
			}

			tempPath := make([]int, len(path), 2*len(path))
			copy(tempPath, path)

			tempPath = append(tempPath, input[i])
			result = append(result, tempPath)

			used[i] = true
			path = append(path, input[i])
			dfs(input, used, i+1)
			path = path[:len(path)-1]
			used[i] = false
		}

	}

	dfs(nums, make([]bool, len(nums), 2*len(nums)), 0)
	return result
}

// @lc code=end

