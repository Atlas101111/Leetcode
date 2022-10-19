/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 *
 * https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (53.00%)
 * Likes:    617
 * Dislikes: 0
 * Total Accepted:    227.5K
 * Total Submissions: 430.6K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * 根据 逆波兰表示法，求表达式的值。
 *
 * 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
 *
 * 注意 两个整数之间的除法只保留整数部分。
 *
 * 可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tokens = ["2","1","+","3","*"]
 * 输出：9
 * 解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
 *
 *
 * 示例 2：
 *
 *
 * 输入：tokens = ["4","13","5","/","+"]
 * 输出：6
 * 解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
 *
 *
 * 示例 3：
 *
 *
 * 输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
 * 输出：22
 * 解释：该算式转化为常见的中缀算术表达式为：
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= tokens.length <= 10^4
 * tokens[i] 是一个算符（"+"、"-"、"*" 或 "/"），或是在范围 [-200, 200] 内的一个整数
 *
 *
 *
 *
 * 逆波兰表达式：
 *
 * 逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
 *
 *
 * 平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
 * 该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
 *
 *
 * 逆波兰表达式主要有以下两个优点：
 *
 *
 * 去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
 * 适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
 *
 *
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, item := range tokens {
		if x, ok := tryParseToDigit(item); ok {
			stack = append(stack, x)
			continue
		}
		operationResult := 0
		second, stack := pop(stack)
		first, stack := pop(stack)
		if item == "*" {
			operationResult = first * second
		} else if item == "/" {
			operationResult = first / second
		} else if item == "+" {
			operationResult = first + second
		} else if item == "-" {
			operationResult = first - second
		}
		stack = push(stack, operationResult)

	}
	return stack[len(stack)-1]
}

func push(stack []int, item int) []int {
	stack = append(stack, item)
	return stack
}

func pop(stack []int) (int, []int) {
	top := stack[len(stack)-1]
	newStack := stack[:len(stack)-1]
	return top, newStack
}

func tryParseToDigit(input string) (int, bool) {
	result := -1
	ok := false
	parse := int(input[0] - '0')
	if parse >= 0 && parse < 10 {
		result = parse
		ok = true
	}
	return result, ok
}

// @lc code=end

