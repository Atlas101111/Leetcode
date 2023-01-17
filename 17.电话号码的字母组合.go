/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (57.99%)
 * Likes:    2268
 * Dislikes: 0
 * Total Accepted:    622K
 * Total Submissions: 1.1M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	path := make([]string, 0)
	hashMap := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return result
	}
	num, _ := strconv.Atoi(digits)

	var dfs func(input int)
	dfs = func(input int) {
		if input == 0 {
			temp := strings.Join(path[:], "")
			result = append(result, temp)
			return
		}

		firstDigit, remains := getFirstDigitAndRemains(input)
		fmt.Printf("%d %d", firstDigit, remains)
		possibleChars := hashMap[firstDigit]

		for _, char := range possibleChars {
			path = append(path, char)
			dfs(remains)
			path = path[:len(path)-1]
		}
	}
	dfs(num)
	return result
}

func getFirstDigitAndRemains(input int) (int, int) {
	result := 0
	count := 0
	remains := input
	for ; input > 0; input /= 10 {
		result = input % 10
		count += 1
	}
	if count == 1 {
		return result, 0
	}
	fmt.Printf("remains before %d\n", remains)
	remains -= result * int(math.Pow(10, float64(count-1)))
	fmt.Printf("remains after %d\n", remains)
	return result, remains
}

// @lc code=end

