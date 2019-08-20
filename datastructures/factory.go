package datastructures

// Factory is a factory object for creating instances of data structures
type Factory struct{}

// TrieNode creates a new Trie node
func (f *Factory) TrieNode() *TrieNode {
	return newTrieNode()
}
