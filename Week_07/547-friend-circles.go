package Week_07

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	parent := make([]int, n) //对称矩阵
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && i != j {
				union(parent, i, j)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}
	return count
}

func find(parent []int, i int) int {
	root := i
	for parent[root] != root {
		root = parent[root]
	}
	for parent[i] != i { //路径压缩
		i, parent[i] = parent[i], root
	}
	return i
}

func union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	if xSet != ySet {
		parent[xSet] = ySet
	}
}
