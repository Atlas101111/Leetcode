/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (71.24%)
 * Likes:    955
 * Dislikes: 0
 * Total Accepted:    757.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [1,null,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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
// package leetcode

// 1.递归法
// func preorderTraversal(root *TreeNode) []int {
// 	result := make([]int, 0)
// 	if root == nil {
// 		return result
// 	}
// 	result = append(result, root.Val)
// 	result = append(result, preorderTraversal(root.Left)...)
// 	result = append(result, preorderTraversal(root.Right)...)
// 	return result
// }

// 迭代法
func preorderTraversal(root *TreeNode) []int {
	var (
		stack  = make([]*TreeNode, 0)
		result = make([]int, 0)
	)
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		l := len(stack)
		top := stack[l-1]
		stack = stack[:l-1]

		result = append(result, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return result
}

// @lc code=end

