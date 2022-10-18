/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 *
 * https://leetcode.cn/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (68.90%)
 * Likes:    765
 * Dislikes: 0
 * Total Accepted:    291.7K
 * Total Submissions: 423.5K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
 *
 * 实现 MyQueue 类：
 *
 *
 * void push(int x) 将元素 x 推到队列的末尾
 * int pop() 从队列的开头移除并返回元素
 * int peek() 返回队列开头的元素
 * boolean empty() 如果队列为空，返回 true ；否则，返回 false
 *
 *
 * 说明：
 *
 *
 * 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty
 * 操作是合法的。
 * 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["MyQueue", "push", "push", "peek", "pop", "empty"]
 * [[], [1], [2], [], [], []]
 * 输出：
 * [null, null, null, 1, 1, false]
 *
 * 解释：
 * MyQueue myQueue = new MyQueue();
 * myQueue.push(1); // queue is: [1]
 * myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
 * myQueue.peek(); // return 1
 * myQueue.pop(); // return 1, queue is [2]
 * myQueue.empty(); // return false
 *
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= x <= 9
 * 最多调用 100 次 push、pop、peek 和 empty
 * 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
 *
 *
 */

// @lc code=start
type MyQueue struct {
	InputStack  []int
	OutputStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		InputStack:  make([]int, 0),
		OutputStack: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.InputStack = append(this.InputStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutputStack) == 0 {
		for i := len(this.InputStack) - 1; i >= 0; i -= 1 {
			this.OutputStack = append(this.OutputStack, this.InputStack[i])
		}
		this.InputStack = make([]int, 0)
	}
	result := this.OutputStack[len(this.OutputStack)-1]
	this.OutputStack = remove(this.OutputStack, len(this.OutputStack)-1)
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.OutputStack) == 0 {
		for i := len(this.InputStack) - 1; i >= 0; i -= 1 {
			this.OutputStack = append(this.OutputStack, this.InputStack[i])
		}
		this.InputStack = make([]int, 0)
	}
	return this.OutputStack[len(this.OutputStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.OutputStack) == 0 && len(this.InputStack) == 0 {
		return true
	}
	return false
}

func remove(list []int, index int) []int {
	return append(list[:index], list[index+1:]...)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

