package 数组

/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//对每个字符排序 记录位置 然后就简单了 本来以为很麻烦的 有简洁的一直不敢下手写
//没想到这个方法和官方方法一样  但是我写的太不简洁 官方go方法很简洁
func groupAnagrams(strs []string) [][]string {
	copyStrs := make([]string, len(strs))
	var result [][]string
	m := make(map[string][]int)
	copy(copyStrs, strs)
	for i := 0; i < len(strs); i++ {
		strs[i] = sortStr(strs[i])
		if _, ok := m[strs[i]]; ok {
			m[strs[i]] = append(m[strs[i]], i)
		} else {
			m[strs[i]] = []int{i}
		}
	}

	for _, v := range m {
		temp := make([]string, len(v))
		for i := 0; i < len(v); i++ {
			temp[i] = copyStrs[v[i]]
		}
		result = append(result, temp)
	}
	return result
}

func sortStr(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		b = append(b, s[i])
	}

	for i := 0; i < len(b)-1; i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	str := ""
	for i := 0; i < len(b); i++ {
		str += string(b[i])
	}
	return str
}
