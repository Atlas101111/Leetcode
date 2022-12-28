/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode.cn/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (57.35%)
 * Likes:    1210
 * Dislikes: 0
 * Total Accepted:    449.2K
 * Total Submissions: 782.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4
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
import "math"

func isBalanced(root *TreeNode) bool {
	// 这个题一开始想到了肯定是递归后序遍历，但是一开始没想明白递归函数的返回
	// 单纯返回左右子树的高度，我们是没有办法完全根据左右子树的高度的来判断整个树是否平衡
	// 如果返回左右子树是否平衡，我们也没有办法得出结论，因为左子树平衡+右子树平衡，并不能推导出整个树平衡这个结论。但是如果有一个子树不平衡，那么整个树一定不平衡
	// 所以这里用了个小技巧，可能是平常工程中不经常这样使用所以没想到，就是用一个-1来标记子树不平衡

	if root == nil {
		return true
	}

	if getDepth(root) != -1 {
		return true
	}

	return false
}

func getDepth(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftHeight, rightHeight := 0, 0

	if root.Left != nil {
		leftHeight = getDepth(root.Left)
	}
	if root.Right != nil {
		rightHeight = getDepth(root.Right)
	}

	if rightHeight == -1 || leftHeight == -1 {
		return -1
	}

	if math.Abs(float64(rightHeight-leftHeight)) > 1 {
		return -1
	}

	if rightHeight > leftHeight {
		return rightHeight + 1
	}
	return leftHeight + 1

}

// @lc code=end

