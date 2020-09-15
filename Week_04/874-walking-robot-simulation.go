package Week_04

type Point struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
	//阻碍点集合
	obMap := make(map[Point]bool)
	for _, ob := range obstacles {
		p := Point{
			x: ob[0],
			y: ob[1],
		}
		obMap[p] = true
	}
	//各个方向的加值
	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d, x, y, res := 0, 0, 0, 0
	for _, c := range commands {
		if c == -1 {
			//右转90
			d++
			if d == 4 { //转了一圈
				d = 0
			}
		} else if c == -2 {
			//左转90
			d--
			if d == -1 { //转了一圈
				d = 3
			}
		} else {
			for ; c > 0; c-- {
				p := Point{
					x: x + move[d][0],
					y: y + move[d][1],
				}
				//检查是否被挡住了
				if _, ok := obMap[p]; !ok {
					x += move[d][0]
					y += move[d][1]
				}
			}
		}
		res = max(res, x*x+y*y)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
