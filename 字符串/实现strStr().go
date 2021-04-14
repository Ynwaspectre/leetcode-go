package 字符串

/**
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0


提示：

0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//本来准备用内置函数的 想想还是自己写个 利用了go的切片,类似于滑动窗口
func strStr(haystack string, needle string) int {
	//haystack的指针
	if len(haystack) < len(needle) {
		return -1
	}
	if len(haystack) == 0 || len(needle) == 0 {
		return 0
	}
	//从左往右  遇到 和 needle首字母相同的就
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
