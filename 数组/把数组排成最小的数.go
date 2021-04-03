package 数组

import (
	"strconv"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。



示例 1:

输入: [10,2]
输出: "102"
示例 2:

输入: [3,30,34,5,9]
输出: "3033459"


提示:

0 < nums.length <= 100
说明:

输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//这道题想了我很久,但是大致思路在最早就确定了，其实就是重新定义一个大小对比方法，不是单纯的数字对比,
//但是在对比方法上想了很久，总是在各个测试用例之间顾此失彼，就在准备看题解的时候,加12 和121 随手写了下
//发现在对比这2个数字的时候 其实就是把12121和12112进行对比嘛 就是12+121 和 121+12 字符串的相加
// 如果12+121>121+12 则 我们就算 12大于121就ok了,排序部分自己写了个比较排序 性能差点
func minNumber(nums []int) string {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if bigger(nums[i], nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	var s string
	for i := 0; i < len(nums); i++ {
		s += strconv.Itoa(nums[i])
	}
	return s
}

func bigger(x, y int) bool {
	return strconv.Itoa(x)+strconv.Itoa(y) > strconv.Itoa(y)+strconv.Itoa(x)
}
