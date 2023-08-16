package trie_node

// 以26个字母为前提测试
type trieNode struct {
	nexts   [26]*trieNode
	passCnt int
	end     bool
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{root: &trieNode{}}
}
