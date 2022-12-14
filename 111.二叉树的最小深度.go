/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (51.40%)
 * Likes:    898
 * Dislikes: 0
 * Total Accepted:    492.3K
 * Total Submissions: 956.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000
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

// 迭代法
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	dep := 0
// 	for len(queue) > 0 {
// 		l := len(queue)
// 		dep += 1

// 		for x := 0; x < l; x += 1 {
// 			front := queue[0]
// 			queue = queue[1:]
// 			if front.Left == nil && front.Right == nil {
// 				return dep
// 			}

// 			if front.Left != nil {
// 				queue = append(queue, front.Left)
// 			}
// 			if front.Right != nil {
// 				queue = append(queue, front.Right)
// 			}
// 		}

// 	}
// 	return dep
// }

// 递归法
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftHeight := -1
	if root.Left != nil {
		leftHeight = minDepth(root.Left)
	}

	rightHeight := -1
	if root.Right != nil {
		rightHeight = minDepth(root.Right)
	}

	if leftHeight == -1 {
		return 1 + rightHeight
	} else if rightHeight == -1 {
		return 1 + leftHeight
	} else if leftHeight > rightHeight {
		return 1 + rightHeight
	} else {
		return 1 + leftHeight
	}

}

// @lc code=end

