package main

import (
	"fmt"
)

func constructArr(a []int) []int {
	b := make([]int, len(a))
	b[0] = 1
	//for i := 1; i < len(a); i++ {
	//	b[i] = b[i-1] * a[i]
	//}
	temp := 1
	for i := len(a) - 1; i >= 0; i-- {
		b[i] = b[i] * temp
		temp *= a[i]
	}

	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(nums))

}
