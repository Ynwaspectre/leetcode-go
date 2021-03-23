package main

import (
	"fmt"
)

func constructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	b := make([]int, len(a)) //构建和a长度一样的b切片
	b[0] = 1                 //b[i] = b[i-1] * a[i] 当i=1的时候b[0]没值 所以取b[0]=1 因为第一个元素是
	///最左边，在左边没有值了 积就是1
	//对下三角进行遍历
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	temp := 1
	for i := len(a) - 2; i >= -1; i-- {
		b[i+1] *= temp //两次的值相乘 得最后的
		temp *= a[i+1]
	}
	return b
}

func main() {
	nums := []int{}
	fmt.Println(constructArr(nums))

}
