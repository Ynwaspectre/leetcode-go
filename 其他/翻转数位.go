package 其他

/**
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。

示例 1：

输入: num = 1775(11011101111)
输出: 8
示例 2：

输入: num = 7(0111)
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-bits-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//简单题竟然 但是我想不出来 本来还打算一天3道简单一道中等的，这种ac率40%的简单题感觉和中等题差不多
//2进制的题目也头疼 不知道为什么不直接出成字符串里1来计算的
//思路看的题解 综合了各个方面选了个容易懂的 稍微修改了 变成自己懂的方式了
func reverseBits(num int) int {
	max := 0    //记录当前最大  最大值等于 遇到0时前面累计的count+1的值加上后来的count的值
	count := 0  //当前正在累计1中 这个count就是当前累计1的值
	insert := 0 //遇到0后 让此值为count+1,count为0  如果下一个还是0 这个值就是1,因为0可以变为1
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			count++
		} else {
			insert = count + 1
			count = 0
		}
		if insert+count > max {
			max = insert + count
		}
		num >>= 1
	}
	return max
}
