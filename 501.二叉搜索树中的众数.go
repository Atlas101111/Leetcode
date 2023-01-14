/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (54.31%)
 * Likes:    567
 * Dislikes: 0
 * Total Accepted:    134.7K
 * Total Submissions: 248.1K
 * Testcase Example:  '[1,null,2,2]'
 *``
 * 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。
 *
 * 如果树中有不止一个众数，可以按 任意顺序 返回。
 *
 * 假定 BST 满足如下定义：
 *
 *
 * 结点左子树中所含节点的值 小于等于 当前节点的值
 * 结点右子树中所含节点的值 大于等于 当前节点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [1, 10^4] 内
 * -10^5 <= Node.val <= 10^5
 *
 *
 *
 *
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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
func findMode(root *TreeNode) []int {
	var pre *TreeNode
	maxCount, count := 0, 0
	result := make([]int, 0)

	var travel func(node *TreeNode)

	travel = func(node *TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)

		if pre == nil {
			count = 1
		} else if pre.Val == node.Val {
			count += 1
		} else {
			count = 1
		}

		if count > maxCount {
			result = make([]int, 0)
			result = append(result, node.Val)
			maxCount = count
		} else if count == maxCount {
			result = append(result, node.Val)
		}
		pre = node

		travel(node.Right)
	}

	travel(root)

	return result
}

// @lc code=end

