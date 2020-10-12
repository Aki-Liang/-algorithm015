package Week_07

type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); i++ {
			dfsTrie(i, j, board, root, &result)
		}
	}
	return result
}

func dfsTrie(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		//访问过的节点或者字典中没有
		return
	}
	node = node.children[c-'a']
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	dfsTrie(i+1, j, board, node, result)
	dfsTrie(i, j+1, board, node, result)
	dfsTrie(i-1, j, board, node, result)
	dfsTrie(i, j-1, board, node, result)
	board[i][j] = c
}
