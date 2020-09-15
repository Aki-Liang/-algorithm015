package Week_04

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		//炸了
		board[click[0]][click[1]] = 'X'
		return board
	}

	dfsBoard(board, click[0], click[1])
	return board
}

//围绕当前方块的8个相邻方块的位移向量
var v = [][]int{{1, -1}, {1, 0}, {1, 1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}}

func dfsBoard(board [][]byte, x, y int) {
	if !inboard(board, x, y) {
		return
	}
	count := 0
	if board[x][y] == 'E' {
		for i := 0; i < 8; i++ {
			//计算周围8个块里有没有雷
			if inboard(board, x+v[i][0], y+v[i][1]) && board[x+v[i][0]][y+v[i][1]] == 'M' {
				count++
			}
		}
		if count > 0 {
			//有雷显示数字
			board[x][y] = '0' + byte(count)
		} else {
			//没雷显示blank，同时扩散出去
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				dfsBoard(board, x+v[i][0], y+v[i][1])
			}
		}
	}
}

func inboard(board [][]byte, x, y int) bool {
	return (x >= 0 && x < len(board) && y >= 0 && y < len(board[0]))
}
