package 数组

/**
机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：

-2 ：向左转 90 度
-1 ：向右转 90 度
1 <= x <= 9 ：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。

机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。

返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）


注意：

北表示 +Y 方向。
东表示 +X 方向。
南表示 -Y 方向。
西表示 -X 方向。


示例 1：

输入：commands = [4,-1,3], obstacles = []
输出：25
解释：
机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 3 个单位，到达 (3, 4)
距离原点最远的是 (3, 4) ，距离为 32 + 42 = 25
示例 2：

输入：commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出：65
解释：机器人开始位于 (0, 0)：
1. 向北移动 4 个单位，到达 (0, 4)
2. 右转
3. 向东移动 1 个单位，然后被位于 (2, 4) 的障碍物阻挡，机器人停在 (1, 4)
4. 左转
5. 向北走 4 个单位，到达 (1, 8)
距离原点最远的是 (1, 8) ，距离为 12 + 82 = 65


提示：

1 <= commands.length <= 104
commands[i] is one of the values in the list [-2,-1,1,2,3,4,5,6,7,8,9].
0 <= obstacles.length <= 104
-3 * 104 <= xi, yi <= 3 * 104
答案保证小于 231

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/walking-robot-simulation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func robotSim(commands []int, obstacles [][]int) int {
	direction := 1 //默认向北
	x, y := 0, 0   //坐标
	flag := 0      //当前是不是遇到障碍物了
	//先标记点
	var nextZhangai []int
	for _, v := range commands {
		if len(obstacles) > 0 {
			nextZhangai = obstacles[0] //最近一个障碍
		}
		if v > 0 && flag == 0 {
			//步数
			if direction == 1 {
				//往北走 x不变 y++
				if len(nextZhangai) > 0 && x == nextZhangai[0] && y+v >= nextZhangai[1] && nextZhangai[1] > y {
					//遇到了 x不变 y为障碍物前面一个
					y = nextZhangai[1] - 1
					flag = 1
				} else {
					y += v
				}
			} else if direction == 2 {
				//向东走 x-- y不变
				if len(nextZhangai) > 0 && y == nextZhangai[1] && x-v <= nextZhangai[0] && x > nextZhangai[0] {
					x = nextZhangai[0] + 1
					flag = 1
				} else {
					x -= v
				}
			} else if direction == 3 {
				//向南走 x不变 y--
				if len(nextZhangai) > 0 && x == nextZhangai[0] && y-v <= nextZhangai[1] && y > nextZhangai[1] {
					y = nextZhangai[1] + 1
					flag = 1
				} else {
					y -= v
				}
			} else {
				//向西走  x++ y不变
				if len(nextZhangai) > 0 && y == nextZhangai[1] && x+v >= nextZhangai[0] && x < nextZhangai[0] {
					x = nextZhangai[0] - 1
					flag = 1
				} else {
					x += v
				}
			}
		} else {
			//改变方向
			direction = getDirection(v, direction)
			if flag == 1 {
				flag = 0
				obstacles = obstacles[1:]
			}
		}

	}
	return x*x + y*y
}

/**
表明方向  北 1  东2  南3  西 4
*/
func getDirection(command int, direction int) int {
	if command == -2 {
		if direction == 1 {
			return 4
		} else if direction == 2 {
			return 1
		} else if direction == 3 {
			return 2
		} else {
			return 3
		}

	} else if command == -1 {
		if direction == 1 {
			return 2
		} else if direction == 2 {
			return 3
		} else if direction == 3 {
			return 4
		} else {
			return 1
		}
	}
	return direction
}
