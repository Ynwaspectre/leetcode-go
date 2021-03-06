package 其他

/**
给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
说明: 你算法的时间复杂度应为 O(log n) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//只能想到算出阶乘然后计算0的个数 这题实际是数学思想 一个0 是2x5 才能构成,整个阶乘中2的个数肯定大于5的个数，
//所以只要算5的个数就可以了 5的阶乘5的个数是1,10的阶乘 5的个数是2  5,5x2这2个  所以n的阶乘里面因子5的个数
//等于n不停的除以5的结果加起来 我肯定想不出来。。数学早忘了
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}
