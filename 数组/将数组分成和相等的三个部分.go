package 数组

/**
给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。

形式上，如果可以找出索引 i+1 < j 且满足 A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1] 就可以将数组三等分。



示例 1：

输入：[0,2,1,-6,6,-7,9,1,2,0,1]
输出：true
解释：0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
示例 2：

输入：[0,2,1,-6,6,7,9,-1,2,0,1]
输出：false
示例 3：

输入：[3,3,6,5,-2,2,5,1,-9,4]
输出：true
解释：3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4


提示：

3 <= A.length <= 50000
-10^4 <= A[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func canThreePartsEqualSum(arr []int) bool {
	totalSum := 0
	for _, v := range arr {
		totalSum += v
	}
	if totalSum%3 != 0 {
		return false
	}
	evePart := totalSum / 3 //每部分的值
	//感觉还是找点符合思维一点
	sum := 0
	i := 0
	n := len(arr)
	for i < n-2 {
		sum += arr[i]
		i++
		if sum == evePart {
			break
		}
	}
	if sum != evePart {
		return false
	}
	j := i
	sum = 0
	for j < n-1 {
		sum += arr[j]
		if sum == evePart {
			break
		}
		j++
	}
	return sum == evePart
}

//func canThreePartsEqualSum(arr []int) bool {
//	totalSum := 0
//	for _, v := range arr {
//		totalSum += v
//	}
//	if totalSum%3 != 0 {
//		return false
//	}
//	evePart := totalSum / 3 //每部分的值
//	//循环 找到3部分
//	sum, count := 0, 0
//	for i := 0; i < len(arr); i++ {
//		sum += arr[i]
//		if sum==evePart {
//			count++
//			sum=0
//		}
//	}
//	//但是这个针对0 0 0 0 这种无法解决 最后必须大于等于3才行,但是为什么大于等于3感觉不太懂
//	return count>=3
//}
