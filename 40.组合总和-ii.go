/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (60.26%)
 * Likes:    1206
 * Dislikes: 0
 * Total Accepted:    383.5K
 * Total Submissions: 636.7K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	result := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(startIndex, target int, used []bool)
	dfs = func(startIndex, target int, used []bool) {
		if target == 0 {
			temp := make([]int, len(path), len(path)*2)
			copy(temp, path)
			result = append(result, temp)
		}

		for x := startIndex; x < len(candidates) && candidates[x] <= target; x += 1 {

			// 这里used[x-1]是要判断为false的时候才剪枝
			// 因为等于false才说明当前的path数组中不含下标x-1的那个元素，说明当前x和x-1的递归深度是一样的
			if x > 0 && used[x-1] == false && candidates[x] == candidates[x-1] {
				continue
			}
			path = append(path, candidates[x])
			used[x] = true
			dfs(x+1, target-candidates[x], used)
			used[x] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0, target, make([]bool, len(candidates), 2*len(candidates)))
	return result
}

// @lc code=end

