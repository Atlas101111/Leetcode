/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (77.10%)
 * Likes:    1448
 * Dislikes: 0
 * Total Accepted:    904.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
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
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	var (
// 		queue = make([]*TreeNode, 0)
// 		depth = 0
// 	)
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		l := len(queue)
// 		depth += 1

// 		for x := 0; x < l; x += 1 {
// 			front := queue[0]
// 			queue = queue[1:]

// 			if front.Left != nil {
// 				queue = append(queue, front.Left)
// 			}
// 			if front.Right != nil {
// 				queue = append(queue, front.Right)
// 			}
// 		}

// 	}
// 	return depth

// }

// 递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// @lc code=end

