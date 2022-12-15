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

// 迭代写法,统一写法,前序遍历
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
			// 为啥每次都是在中节点后面插入一个nil呢？是因为我们事实上是在用nil标记我们走过的路径，中节点是我们肯定踏足过的节点，所以要加标记

			// 这样做的原因是，在后序遍历和中序遍历的时候，当我们访问到某个中节点A，并不代表我们要立刻输出A的值或者对A节点做处理，有可能只是路过，
			// 并不需要A节点的值，只是为了获取节点A的左右子节点而已，按照遍历顺序还轮不到A，所以用nil来标记路过但没访问过值的节点

		} else {
			ptr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ptr.Left, ptr.Right = ptr.Right, ptr.Left
		}
	}
	return root
}

// @lc code=end

