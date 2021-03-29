package 数组

import (
	"math"
)

/**
给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。



进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？



示例 1：

输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。
示例 2：

输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。
示例 3：

输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。


提示：

n == nums.length
1 <= n <= 104
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/132-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//草稿本
//双指针。
//构造map 存储 数字=>索引【】
//i 的范围是0~len(nums)-3
//遍历i的话 问题就转变成 从数组里找2个数j和k 满足j<k 而且nums[j]>nums[k] nums[k]>nums[i]  好的 果断O(n^3)超时
//func find132pattern(nums []int) bool {
//	var temp int
//	for i := 0; i < len(nums); i++ {
//		temp = nums[i]
//		if find(temp,nums[i+1:]){
//			return true
//		}
//	}
//	return false
//}
//
////从nums里找2个数  a和b  a>b b>tem 而且a的索引在b前面
//func find(temp int, nums []int) bool {
//	for i := 0; i < len(nums); i++ {
//		for j:=i+1;j<len(nums);j++{
//			if (nums[j]<nums[i]) &&(nums[j]>temp){
//				return true
//			}
//		}
//	}
//	return false
//}

//思路是从右往左遍历 先设置一个变量min为最小值 含义为小于3的最大值 即为当前可以用的2里面的最大值
//为什么要最大值呢 因为如果左边1的元素得比2小 所以我们要尽可能让2是小于3的最大值
// stack存储的是啥呢 存储的是一堆能满足的2 不过当新遍历的元素>stack中的最大元素得时候，
//这个新遍历的元素就可以当做3 而把stack中的最大元素赋值给之前那个变量min
//-1,3,2,0
func find132pattern(nums []int) bool {
	length := len(nums)
	min := math.MinInt64
	stack := []int{nums[length-1]}
	for i := length - 2; i >= 0; i-- {
		//当新遍历的元素 小于我们存储的最大的2的时候 代表有一个132模式 ，这个1,3,2模式就是nums[i],stack的最大元素,min值
		if nums[i] < min {
			return true
		}
		//找stack里的最大值 看看是否比nums[i]小
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			min = stack[len(stack)-1]    //比nums[i]小的数 赋值给了min
			stack = stack[:len(stack)-1] //丢到上一步的数
		}
		if nums[i] > min {
			stack = append(stack, nums[i])
		}
	}
	return false
}

//栈中最大的元素即当前遍历过的元素里最大的一个
//总思路 就是从右到左遍历  维护 一个最大的2 和一个单调栈 其中栈顶是3比2大，如果接下来的数比2小 那就ok；否则把栈中小于这个数的值都拿出来，
//最后一个赋值给2 ,然后把3放进去;
//太难消化了这题  看别人的思路去想去消化 和自己想的要消耗的脑力差太多了
