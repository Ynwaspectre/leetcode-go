package 数组

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。



示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//刚开始思路想对了 就是明确方向 然后调整i和j的增减，用总数量来判断循环的截止，但是没写下去，晚上看了题解发现官方方法1和我的思路一样
//但是粗略一看没懂 有个什么%4没了解啥意思，后面自己按自己写法来，知道%4是啥意思了 类似于一个开关 只有1和2两个值,取反的话可以用3-这个值，
//因为这个题目中要变方向,自己随便定义了方向 然后写了个改变方向的方法,还有一个判断是否访问过的 本来没看官方的时候也想到了要保存，想用map的
//还是官方的简单点 ,代码写的很乱 还有很多可以优化的

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	row := len(matrix) - 1
	col := len(matrix[0]) - 1

	//新见一个矩阵代表是否走过了
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		v := make([]bool, len(matrix[0]))
		visited[i] = v
	}

	totalNum := len(matrix) * len(matrix[0])
	//存储遍历过的坐标
	//记录当前的位置 如果是往右 要变化只能往下
	direction := 0 //0 向右  1 向下 2 向左 3向上
	var store []int
	i := 0
	j := 0
	for totalNum > 0 {
		if visited[i][j] { //访问过 修改方向
			if direction == 0 {
				j--
				i++
			} else if direction == 1 {
				i--
				j--
			} else if direction == 2 {
				j++
				i--
			} else {
				i++
				j++
			}
			direction = changeDir(direction)
			continue
		} else {
			store = append(store, matrix[i][j]) //存起来
			visited[i][j] = true
			//判断是否到角了
			if direction == 0 {
				if j == col {
					direction = changeDir(direction)
					i++
				} else {
					j++
				}
			} else if direction == 1 {
				if i == row {
					direction = changeDir(direction)
					j--
				} else {
					i++
				}
			} else if direction == 2 {
				if j == 0 {
					direction = changeDir(direction)
					i--
				} else {
					j--
				}
			} else {
				if i == 0 {
					direction = changeDir(direction)
					j++
				} else {
					i--
				}
			}
		}
		totalNum--
	}
	return store
}

func changeDir(direction int) int {
	if direction == 3 {
		direction = 0
	} else {
		direction++
	}
	return direction
}
