package 其他

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成


进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//dfs 需要一个额外空间来标明有没访问过  只想到这么多 其余看的题解
//细节方面太难处理了 第70行这个我比较难理解 应该是对于35~41的单词循环
//使用的visited是同一个 如果在某个方向寻找下去不成功的话 则应该把标记过的重新置为false
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		temp := make([]bool, len(board[0]))
		visited[i] = temp
	}

	//遍历每个元素
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, visited, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//从index开始找
func dfs(board [][]byte, visited [][]bool, word string, i int, j int, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	//查看相邻位置是否可以
	visited[i][j] = true
	//寻找相邻位置
	directs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := false
	for _, v := range directs {
		newI := i + v[0]
		newJ := j + v[1]
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
			if !visited[newI][newJ] {
				if dfs(board, visited, word, newI, newJ, k+1) {
					result = true
					break
				}
			}
		}
	}
	visited[i][j] = false
	return result
}
