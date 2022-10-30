/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode.cn/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (50.03%)
 * Likes:    1920
 * Dislikes: 0
 * Total Accepted:    369.3K
 * Total Submissions: 738.5K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回 滑动窗口中的最大值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */

// @lc code=start
type MonoDecQueue struct {
	queue []int
}

func NewMonoDecQueue() *MonoDecQueue {
	return &MonoDecQueue{
		queue: make([]int, 0),
	}
}

func (q *MonoDecQueue) IsEmpty() bool {
	if len(q.queue) == 0 {
		return true
	}
	return false
}

func (q *MonoDecQueue) Front() int {
	if q.IsEmpty() {
		return 0
	}
	return q.queue[0]
}

func (q *MonoDecQueue) Tail() int {
	if q.IsEmpty() {
		return 0
	}
	return q.queue[len(q.queue)-1]
}

func (q *MonoDecQueue) Push(x int) {
	for !q.IsEmpty() && q.Tail() < x {
		q.PopBackIfNeed(q.Tail())
	}
	q.queue = append(q.queue, x)
}

func (q *MonoDecQueue) PopFrontIfNeed(x int) {
	if q.Front() == x {
		q.queue = q.queue[1:]
	}
}

func (q *MonoDecQueue) PopBackIfNeed(x int) {
	if q.Tail() == x {
		q.queue = q.queue[:len(q.queue)-1]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMonoDecQueue()
	result := make([]int, 0)
	for i := 0; i < k && i < len(nums); i += 1 {
		queue.Push(nums[i])
	}
	result = append(result, queue.Front())

	if len(nums) <= k {
		return result
	}

	for j := 1; j+k-1 <= len(nums)-1; j += 1 {

		queue.PopFrontIfNeed(nums[j-1])
		queue.Push(nums[j+k-1])
		result = append(result, queue.Front())
	}
	return result
}

// @lc code=end

