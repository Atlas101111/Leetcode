/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode.cn/problems/path-sum/description/
 *
 * algorithms
 * Easy (53.47%)
 * Likes:    1074
 * Dislikes: 0
 * Total Accepted:    487.3K
 * Total Submissions: 910.5K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点
 * 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
 *
 * 叶子节点 是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
 * 输出：true
 * 解释：等于目标和的根节点到叶节点路径如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,3], targetSum = 5
 * 输出：false
 * 解释：树中存在两条根节点到叶子节点的路径：
 * (1 --> 2): 和为 3
 * (1 --> 3): 和为 4
 * 不存在 sum = 5 的根节点到叶子节点的路径。
 *
 * 示例 3：
 *
 *
 * 输入：root = [], targetSum = 0
 * 输出：false
 * 解释：由于树是空的，所以不存在根节点到叶子节点的路径。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 5000] 内
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return getResult2(root, targetSum)
}

func getResult(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		// 这里的递归边界只能是到叶子结点，不能是空节点
		// 可以用TestCase：[1,2]\n1  来试一下getResult2这个函数
		// 因为路径的终点一定是一个叶子结点
		return (targetSum - root.Val) == 0
	}

	leftResult, rightResult := false, false

	if root.Left != nil {
		leftResult = getResult(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		rightResult = getResult(root.Right, targetSum-root.Val)
	}

	return leftResult || rightResult

}

func getResult2(root *TreeNode, targetSum int) bool {
	if root == nil {
		if targetSum == 0 {
			return true
		}
		return false
	}

	leftResult, rightResult := false, false
	leftResult = getResult2(root.Left, targetSum-root.Val)
	rightResult = getResult2(root.Right, targetSum-root.Val)

	return leftResult || rightResult
}

// @lc code=end

