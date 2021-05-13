package 其他

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//2=>abc
//3=>def
//4=>ghi
//5=>jkl
//6=>mno
//7=>pqrs
//8=>tuv
//9=>wxyz

//看了下 好像没遇到这种题目类型过 这要全部遍历0(n^4)了 最多是4个数字 最多的情况 3 3 4 4 个字符
// 递归   4个字符的 等于3个字符的答案加上 3个字符的答案每个都与第四个字符组成的那个就是最终答案
// 这题好像不算单个字符 就算digits长度的字符 所以第一次提交失败 第二次ok
func letterCombinations(digits string) []string {
	length := len(digits)

	if length == 0 {
		return []string{}
	}
	if length == 1 {
		return getStr(string(digits[0]))
	}

	a := letterCombinations(digits[0 : length-1])
	var c []string
	for _, v := range a {
		sarr := getStr(string(digits[length-1]))
		for _, v1 := range sarr {
			c = append(c, v+v1)
		}
	}
	//a = append(a, c...)
	return c
}

func getStr(number string) []string {
	switch number {
	case "2":
		return []string{"a", "b", "c"}
	case "3":
		return []string{"d", "e", "f"}
	case "4":
		return []string{"g", "h", "i"}
	case "5":
		return []string{"j", "k", "l"}
	case "6":
		return []string{"m", "n", "o"}
	case "7":
		return []string{"p", "q", "r", "s"}
	case "8":
		return []string{"t", "u", "v"}
	case "9":
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}
