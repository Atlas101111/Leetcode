/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.27%)
 * Likes:    1248
 * Dislikes: 0
 * Total Accepted:    470.4K
 * Total Submissions: 608.7K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	var travel func(tail int, count int, pre []int)
	travel = func(tail int, count int, pre []int) {
		if count == 0 {
			result = append(result, pre)
			return
		}
		// 这里的剪枝条件到底应该是x>=count-1,还是x>count-1?
		// 也可以写成x-1 >= count, 更好理解一些
		nextCount := count - 1
		for x := tail; x > nextCount; x -= 1 {
			// 下面这种写法纯属Golang的大坑
			// next := append(pre, x)
			// 这里有时候会分配新内存，有时候又不会
			// 取决于pre数组的cap，只有cap改变时，next才会深拷贝

			// 下面这种深拷贝也有坑，要保证next在初始化的时候，len大于等于目标数组，cap是目标数组的两倍
			next := make([]int, len(pre), len(pre)*2)
			copy(next, pre)
			next = append(next, x)

			travel(x-1, nextCount, next)
		}

	}

	travel(n, k, make([]int, 0))
	return result
}

// @lc code=end

