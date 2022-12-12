/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 *
 * https://leetcode.cn/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (56.83%)
 * Likes:    507
 * Dislikes: 0
 * Total Accepted:    95K
 * Total Submissions: 167.5K
 * Testcase Example:  '["i","love","leetcode","i","love","coding"]\n2'
 *
 * 给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
 *
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 * ⁠   注意，按字母顺序 "i" 在 "love" 之前。
 *
 *
 * 示例 2：
 *
 *
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
 * k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 * ⁠   出现次数依次为 4, 3, 2 和 1 次。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 1 <= words.length <= 500
 * 1 <= words[i] <= 10
 * words[i] 由小写英文字母组成。
 * k 的取值范围是 [1, 不同 words[i] 的数量]
 *
 *
 *
 *
 * 进阶：尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
 *
 */

// @lc code=start
// package leetcode

import (
	"container/heap"
)

type Node struct {
	Word      string
	Frequency int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	if h[i].Frequency == h[j].Frequency {
		return h[i].Word > h[j].Word
	}
	// 小顶堆
	return h[i].Frequency < h[j].Frequency
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	oldLen := len(old)
	result := old[oldLen-1]
	*h = old[:oldLen-1]
	return result
}

func (h *NodeHeap) Push(node interface{}) {
	*h = append(*h, node.(*Node))
}

func topKFrequent(words []string, k int) []string {
	dict := make(map[string]*Node)
	// On统计单词的频率
	for _, word := range words {
		if val, ok := dict[word]; !ok {
			newNode := &Node{
				Word:      word,
				Frequency: 1,
			}
			dict[word] = newNode
		} else {
			val.Frequency += 1
		}
	}

	var (
		myHeap = NodeHeap{}
		result = make([]string, k)
	)
	heap.Init(&myHeap)

	// 用堆的主要点在这里，维护一个大小为k的小顶堆，将元素按照任意顺序插入堆中
	// 每当堆的大小超过k，就弹出堆顶，即目前最小的元素
	// 插入完成后，剩下的k个元素即为TopK元素
	for _, value := range dict {
		// 这里有一个点，就是一定要用heap.Push,不能用myHeap.Push，要不然就不会走树的修正流程
		heap.Push(&myHeap, value)
		if len(myHeap) > k {
			heap.Pop(&myHeap)
		}
	}
	// for _, node := range myHeap {
	// 	fmt.Printf("word: %s, frq: %v\n", node.Word, node.Frequency)
	// }

	// 必须按照Pop的顺序，逆序输出，才是从大到小的顺序
	for x := 0; x < k; x += 1 {
		node := heap.Pop(&myHeap)
		result[k-1-x] = node.(*Node).Word
	}

	return result
}

// @lc code=end
