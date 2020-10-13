package Week_07

func find(parents []int, i int) int {
	root := i
	for parents[root] != root {
		root = parents[root]
	}

	for parents[i] != i {
		i, parents[i] = parents[i], root
	}

	return i
}

func union(parents []int, i, j int) {
	ri := find(parents, i)
	rj := find(parents, j)
	if ri != rj {
		parents[ri] = rj
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	parents := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		parents[i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}

			pos := i*m + j
			if j+1 < m && grid[i][j+1] == '1' {
				//如果右边是陆地，合并
				union(parents, pos, pos+1)
			}
			if i+1 < n && grid[i+1][j] == '1' {
				//如果下边是陆地，合并
				union(parents, pos, pos+m)

			}
		}
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && parents[i*m+j] == i*m+j {
				islands++
			}
		}
	}

	return islands
}
