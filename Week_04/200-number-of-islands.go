package Week_04

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if '0' == grid[line][column] { //找到陆地
				continue
			}
			count++
			dfs(grid, line, column)
		}
	}

	return count
}

func dfs(grid [][]byte, line, column int) {
	//找到相连的陆地，然后搞沉没
	if line < 0 || line >= len(grid) || column < 0 || column >= len(grid[line]) {
		return
	}
	if grid[line][column] != '1' {
		//没有陆地了
		return
	}
	grid[line][column] = '0'
	dfs(grid, line+1, column)
	dfs(grid, line-1, column)
	dfs(grid, line, column+1)
	dfs(grid, line, column-1)

}
