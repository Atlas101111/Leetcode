/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode.cn/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Hard (47.37%)
 * Likes:    714
 * Dislikes: 0
 * Total Accepted:    76K
 * Total Submissions: 160K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi]
 * 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
 *
 * 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK
 * 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
 *
 *
 * 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
 *
 *
 * 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
 * 输出：["JFK","MUC","LHR","SFO","SJC"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：tickets =
 * [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * tickets[i].length == 2
 * fromi.length == 3
 * toi.length == 3
 * fromi 和 toi 由大写英文字母组成
 * fromi != toi
 *
 *
 */

// @lc code=start
type info struct {
	target string
	used   bool
}
type infos []*info

func (ins infos) Len() int {
	return len(ins)
}

func (ins infos) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
	return
}

func (ins infos) Less(i, j int) bool {
	return ins[i].target < ins[j].target
}

func findItinerary(tickets [][]string) []string {
	path := make([]string, 0)
	path = append(path, "JFK")

	ticketInfos := make(map[string]infos)
	for _, item := range tickets {
		depart := item[0]
		if ticketInfos[depart] == nil {
			ticketInfos[depart] = make(infos, 0)
		}
		ticketInfos[depart] = append(ticketInfos[depart], &info{target: item[1]})
	}
	for k, _ := range ticketInfos {
		sort.Sort(ticketInfos[k])
	}

	var dfs func(unusedCount int) bool
	dfs = func(unusedCount int) bool {
		if unusedCount == 0 {
			return true
		}
		from := path[len(path)-1]
		for _, ticket := range ticketInfos[from] {
			if ticket.used {
				continue
			}

			path = append(path, ticket.target)
			ticket.used = true
			if dfs(unusedCount - 1) {
				return true
			}
			ticket.used = false
			path = path[:len(path)-1]
		}
		return false
	}

	dfs(len(tickets))
	return path
}

// @lc code=end

