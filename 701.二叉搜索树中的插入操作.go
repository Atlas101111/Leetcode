/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
 *
 * https://leetcode.cn/problems/insert-into-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (71.35%)
 * Likes:    417
 * Dislikes: 0
 * Total Accepted:    165.8K
 * Total Submissions: 232.3K
 * Testcase Example:  '[4,2,7,1,3]\n5'
 *
 * 给定二叉搜索树（BST）的根节点 root 和要插入树中的值 value ，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证
 * ，新值和原始二叉搜索树中的任意节点值都不同。
 *
 * 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [4,2,7,1,3], val = 5
 * 输出：[4,2,7,1,3,5]
 * 解释：另一个满足题目要求可以通过的树是：
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [40,20,60,10,30,50,70], val = 25
 * 输出：[40,20,60,10,30,50,70,null,null,25]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
 * 输出：[4,2,7,1,3,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数将在 [0, 10^4]的范围内。
 * -10^8 <= Node.val <= 10^8
 * 所有值 Node.val 是 独一无二 的。
 * -10^8 <= val <= 10^8
 * 保证 val 在原始BST中不存在。
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

// package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func insertIntoBST(root *TreeNode, val int) *TreeNode {
// 	if root == nil {
// 		root = &TreeNode{
// 			Val: val,
// 		}
// 		return root
// 	}

// 	var travel func(node *TreeNode, target int)

// 	travel = func(node *TreeNode, target int) {
// 		if node == nil {
// 			return
// 		}
// 		if target > node.Val {
// 			if node.Right == nil {
// 				node.Right = &TreeNode{
// 					Val: target,
// 				}
// 				return
// 			}
// 			travel(node.Right, target)
// 		} else {
// 			if node.Left == nil {
// 				node.Left = &TreeNode{
// 					Val: target,
// 				}
// 				return
// 			}
// 			travel(node.Left, target)
// 		}
// 	}

// 	travel(root, val)
// 	return root
// }

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

// @lc code=end
