/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode.cn/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (58.44%)
 * Likes:    2216
 * Dislikes: 0
 * Total Accepted:    737.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
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

// 递归
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return compareTree(root.Left, root.Right)
// }

// func compareTree(leftTree *TreeNode, rightTree *TreeNode) bool {
// 	if (leftTree == nil && rightTree != nil) || (leftTree != nil && rightTree == nil) {
// 		return false
// 	}
// 	if leftTree == nil && rightTree == nil {
// 		return true
// 	}
// 	if leftTree.Val != rightTree.Val {
// 		return false
// 	}

// 	leftEql := compareTree(leftTree.Left, rightTree.Right)
// 	rightEql := compareTree(leftTree.Right, rightTree.Left)

// 	if leftEql && rightEql {
// 		return true
// 	}
// 	return false
// }

// 迭代
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var (
		queue = make([]*TreeNode, 0)
	)
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		leftFront := queue[0]
		queue = queue[1:]
		rightTail := queue[0]
		queue = queue[1:]

		if (leftFront == nil && rightTail != nil) || (leftFront != nil && rightTail == nil) {
			return false
		}
		if leftFront == nil || rightTail == nil {
			continue
		}

		if leftFront.Val != rightTail.Val {
			return false
		}

		queue = append(queue, leftFront.Left)
		queue = append(queue, rightTail.Right)
		queue = append(queue, leftFront.Right)
		queue = append(queue, rightTail.Left)
	}
	return true
}

// @lc code=end

