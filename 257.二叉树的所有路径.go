/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode.cn/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (70.49%)
 * Likes:    858
 * Dislikes: 0
 * Total Accepted:    272.1K
 * Total Submissions: 385.7K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,5]
 * 输出：["1->2->5","1->3"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：["1"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 100] 内
 * -100 <= Node.val <= 100
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

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	if root == nil {
		return result
	}

	searchResult := deepFirstSearch(root, strconv.Itoa(root.Val))

	return searchResult

}

func deepFirstSearch(root *TreeNode, last string) []string {
	searchResult := make([]string, 0)
	if root.Left == nil && root.Right == nil {
		// result := []string{last}
		// lastResult = append(lastResult, last)
		searchResult = append(searchResult, last)
		return searchResult
	}

	if root.Left != nil {
		nextLevelResult := deepFirstSearch(root.Left, last+"->"+strconv.Itoa(root.Left.Val))
		searchResult = append(searchResult, nextLevelResult...)

	}

	if root.Right != nil {
		nextLevelResult := deepFirstSearch(root.Right, last+"->"+strconv.Itoa(root.Right.Val))
		searchResult = append(searchResult, nextLevelResult...)
	}

	return searchResult
}

// @lc code=end

