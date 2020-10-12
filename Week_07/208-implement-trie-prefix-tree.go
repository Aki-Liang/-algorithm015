package Week_07

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		if node.next[c-'a'] == nil {
			node.next[c-'a'] = &Trie{}
		}
		node = node.next[c-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		if node = node.next[c-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		if node = node.next[c-'a']; node == nil {
			return false
		}
	}
	return true
}
