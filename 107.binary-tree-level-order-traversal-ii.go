/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Medium (72.07%)
 * Likes:    637
 * Dislikes: 0
 * Total Accepted:    247.9K
 * Total Submissions: 344K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[15,7],[9,20],[3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
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
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
func levelOrderBottom(root *TreeNode) [][]int {
	var (
		queue       = make([]*TreeNode, 0)
		result      = make([][]int, 0)
		resultStack = make([][]int, 0)
	)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]int, 0)
		l := len(queue)
		for x := 0; x < l; x += 1 {
			top := queue[0]
			queue = queue[1:]

			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		resultStack = append(resultStack, temp)
	}

	for len(resultStack) > 0 {
		top := resultStack[len(resultStack)-1]
		resultStack = resultStack[:len(resultStack)-1]
		result = append(result, top)
	}
	return result
}

// @lc code=end

