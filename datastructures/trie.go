package datastructures

// TrieNode represents a node in a trie
type TrieNode struct {
	IsWord   bool
	Children map[rune]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{IsWord: false, Children: map[rune]*TrieNode{}}
}

// Add inserts a word into a trie
func (t *TrieNode) Add(w string) {
	curr := t
	for _, ch := range w {
		if cn, ok := curr.Children[ch]; ok {
			curr = cn
			continue
		}
		curr.Children[ch] = newTrieNode()
		curr = curr.Children[ch]
	}
	curr.IsWord = true
}
