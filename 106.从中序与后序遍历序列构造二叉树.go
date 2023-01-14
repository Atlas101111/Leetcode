/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.46%)
 * Likes:    904
 * Dislikes: 0
 * Total Accepted:    250K
 * Total Submissions: 345.2K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入：inorder = [-1], postorder = [-1]
 * 输出：[-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder 和 postorder 都由 不同 的值组成
 * postorder 中每一个值都在 inorder 中
 * inorder 保证是树的中序遍历
 * postorder 保证是树的后序遍历
 *
 *
 */

// @lc code=start

// package leetcode

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	delimiter := postorder[len(postorder)-1]
	delimiterIndex := 0
	for index, element := range inorder {
		if element == delimiter {
			delimiterIndex = index
			break
		}
	}

	leftInorder := inorder[:delimiterIndex]
	rightInorder := inorder[delimiterIndex+1:]

	postorder = postorder[:len(postorder)-1]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):]

	root := TreeNode{
		Val:   delimiter,
		Left:  buildTree(leftInorder, leftPostorder),
		Right: buildTree(rightInorder, rightPostorder),
	}

	return &root
}

// @lc code=end
