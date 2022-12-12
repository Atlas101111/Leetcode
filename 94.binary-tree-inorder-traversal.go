/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (76.13%)
 * Likes:    1633
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
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
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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

// 迭代法
func inorderTraversal(root *TreeNode) []int {
	var (
		stack  = make([]*TreeNode, 0)
		result = make([]int, 0)
		ptr    = root
	)
	for ptr != nil || len(stack) > 0 {
		if ptr != nil {
			stack = append(stack, ptr)
			ptr = ptr.Left
		} else {
			l := len(stack)
			top := stack[l-1]
			stack = stack[:l-1]

			result = append(result, top.Val)
			ptr = top.Right
		}
	}
	return result
}

// @lc code=end

