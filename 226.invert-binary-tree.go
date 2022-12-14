/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode.cn/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (79.40%)
 * Likes:    1459
 * Dislikes: 0
 * Total Accepted:    583K
 * Total Submissions: 734.3K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [4,2,7,1,3,6,9]
 * 输出：[4,7,2,9,6,3,1]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,3]
 * 输出：[2,3,1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在 [0, 100] 内
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

// 递归写法
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}
// 	root.Left, root.Right = root.Right, root.Left
// 	invertTree(root.Left)
// 	invertTree(root.Right)
// 	return root
// }

// 迭代写法,统一写法
func invertTree(root *TreeNode) *TreeNode {
	var (
		stack = make([]*TreeNode, 0)
		ptr   = root
	)
	if root == nil {
		return root
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		ptr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if ptr != nil {
			if ptr.Right != nil {
				stack = append(stack, ptr.Right)
			}
			if ptr.Left != nil {
				stack = append(stack, ptr.Left)
			}
			stack = append(stack, ptr)
			stack = append(stack, nil)

		} else {
			ptr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ptr.Left, ptr.Right = ptr.Right, ptr.Left
		}
	}
	return root
}

// @lc code=end

