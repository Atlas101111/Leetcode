/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode.cn/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (63.20%)
 * Likes:    882
 * Dislikes: 0
 * Total Accepted:    315K
 * Total Submissions: 498.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 输出：[[5,4,11,2],[5,8,4,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], targetSum = 0
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点总数在范围 [0, 5000] 内
 * -1000
 * -1000
 *
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	result = getResult(root, targetSum, make([]int, 0))

	return result
}

func getResult(root *TreeNode, targetSum int, path []int) [][]int {
	result := make([][]int, 0)
	newPath := append(path, root.Val)
	newSum := targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		if newSum == 0 {
			result = append(result, newPath)
			return result
		}
		return result
	}
	if root.Left != nil {
		leftResult := getResult(root.Left, newSum, newPath)
		if len(leftResult) > 0 {
			result = append(result, leftResult...)
		}
	}
	if root.Right != nil {
		rightResult := getResult(root.Right, newSum, newPath)
		if len(rightResult) > 0 {
			result = append(result, rightResult...)
		}
	}

	return result
}

// @lc code=end

