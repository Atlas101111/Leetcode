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

var (
	result [][]int
	path   []int
	used   []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	result, path = make([][]int, 0), make([]int, 0)
	used = make([]bool, len(candidates))
	// 这里排序是必须的，
	sort.Ints(candidates)
	dfs(candidates, target, 0)
	return result
}

func dfs(candidates []int, target, index int) {
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if i >= 1 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, candidates[i])
		target -= candidates[i]
		used[i] = true
		dfs(candidates, target, i+1)
		path = path[:len(path)-1]
		target += candidates[i]
		used[i] = false
	}
}

// @lc code=end

