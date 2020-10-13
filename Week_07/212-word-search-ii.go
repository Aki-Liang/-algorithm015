package Week_07

type Trie struct {
	word string
	next [26]*Trie
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	root := &Trie{}
	for _, word := range words {
		node := root
		for _, c := range word {
			if node.next[c-'a'] == nil {
				node.next[c-'a'] = &Trie{}
			}
			node = node.next[c-'a']
		}
		node.word = word
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if root.next[board[i][j]-'a'] == nil {
				continue
			}
			wDfs(i, j, board, root, &res)
		}
	}

	return res
}

func wDfs(i, j int, board [][]byte, node *Trie, res *[]string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '#' {
		return
	}

	c := board[i][j]
	if node.next[c-'a'] == nil {
		return
	}
	node = node.next[c-'a']
	if node.word != "" {
		*res = append(*res, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	wDfs(i+1, j, board, node, res)
	wDfs(i-1, j, board, node, res)
	wDfs(i, j+1, board, node, res)
	wDfs(i, j-1, board, node, res)
	board[i][j] = c

}
