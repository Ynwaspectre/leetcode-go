package 数组

/**
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//4次才过。。。纯自己想 在想规律的时候一些操作绕了点
//总的思路就是 先倒叙查 如果能找到后面一个大于前面的 就是 有下一个最大的
//否则进行翻转

//

func nextPermutation(nums []int) {
	flag := false //是否可以找到
	keyIndex := -1
	firstMaxIndex := -1
	//倒叙如果能找到 后面的大于前面的就代表可以找到 下一个最大的
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			flag = true
			keyIndex = i - 1
			break //可以找到
		}
	}
	if !flag {
		swap(nums, 0, len(nums)-1)
	} else {
		//找第一个比keyIndex大的
		for i := len(nums) - 1; i >= 1; i-- {
			if nums[i] > nums[keyIndex] {
				firstMaxIndex = i
				break
			}
		}
		swap(nums, keyIndex+1, len(nums)-1)
		nums[keyIndex], nums[keyIndex+len(nums)-firstMaxIndex] = nums[keyIndex+len(nums)-firstMaxIndex], nums[keyIndex]
	}
}

func swap(nums []int, l, r int) {
	mid := (l + r) / 2
	for i := l; i <= mid; i++ {
		nums[i], nums[l+r-i] = nums[l+r-i], nums[i]
	}
}
