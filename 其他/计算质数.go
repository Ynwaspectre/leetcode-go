package 其他

/**
统计所有小于非负整数 n 的质数的数量。



示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：

输入：n = 0
输出：0
示例 3：

输入：n = 1
输出：0


提示：

0 <= n <= 5 * 106

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-primes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	isp := make([]bool, n)
	for i := 0; i < n; i++ {
		isp[i] = true
	}
	count := 0
	for i := 2; i < n; i++ {
		if isp[i] {
			count++
			if i*i < n {
				for j := i * i; j < n; j += i { //j+=i  i倍的i （i+1）倍的i
					isp[j] = false
				}
			}
		}
	}
	return count
}
