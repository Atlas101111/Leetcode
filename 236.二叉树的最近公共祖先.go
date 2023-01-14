/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (69.52%)
 * Likes:    2092
 * Dislikes: 0
 * Total Accepted:    486.9K
 * Total Submissions: 700.4K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
 * 输出：3
 * 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
 * 输出：5
 * 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], p = 1, q = 2
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [2, 10^5] 内。
 * -10^9
 * 所有 Node.val 互不相同 。
 * p != q
 * p 和 q 均存在于给定的二叉树中。
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 这个题感觉像分治+回溯

	// 我们先定义递归函数的语意，函数接受root，p，q，返回在root为根的树下，p和q的最近公共祖先
	// 当p，q不同时存在时，返回存在的那个p或者q，相当于p或q为自己的公共祖先，根据定义最近公共祖先节点可以为节点本身；
	// 当p和q都不存在，相当于没有公共祖先，则返回nil

	leftResult := lowestCommonAncestor(root.Left, p, q)
	rightResult := lowestCommonAncestor(root.Right, p, q)

	// 这里要先判断，当前的树根是否为p或者q，这样才符合函数语意
	// 况且这里自底向上的访问顺序，其实是后序遍历，用左右子树的结果以及root的结果，共同得出当前子树是否存在最近公共祖先节点
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	if leftResult != nil && rightResult != nil {
		return root
	} else if leftResult == nil && rightResult == nil {
		return nil
	} else {
		// 这里也是，不能因为当前子树不同时存在p和q就返回nil，才符合最近公共祖先节点可以为节点本身的语意，才能把左右子树的结果一层层传递到上面
		if leftResult != nil {
			return leftResult
		} else {
			return rightResult
		}
	}
	return nil

}

// @lc code=end

