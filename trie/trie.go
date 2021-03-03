package trie

type TrieNode struct {
	value    byte
	parent   *TrieNode
	children [26]*TrieNode
}

func NewTrieNode(val byte, par *TrieNode) *TrieNode {
	return &TrieNode{value: val, parent: par}
}

type Trie struct {
	root TrieNode
}

func NewTrie() *Trie {
	rootNode := NewTrieNode(0, nil)
	return &Trie{root: *rootNode}
}

// TODO SB: Probably want to return last node of work
func (t *Trie) InsertWord(word string) {
	RecurseInsert(word, &t.root)
}

func (t *Trie) FindWord(word string) (bool, *TrieNode) {
	asByte := byte(word[0])
	if t.root.children[asByte-97] == nil {
		return false, nil
	}
	return RecurseFind(word[1:], t.root.children[asByte-97])
}

// TODO SB: I think I have my concepts of current node vs place in the word a little mixed up here
func RecurseFind(restOfWord string, current *TrieNode) (bool, *TrieNode) {
	// TODO SB: Can we be sure this guarantees finding the word?
	if 0 == len(restOfWord) {
		return true, current
	}
	asByte := byte(restOfWord[0])
	if current.children[asByte-97] == nil {
		return false, nil
	}
	return RecurseFind(restOfWord[1:], current.children[asByte-97])
}

// TODO SB: Currently this will always return nil i think, probably want to return last node
func RecurseInsert(word string, parent *TrieNode) *TrieNode {
	if 0 == len(word) {
		return nil
	}
	asByte := byte(word[0])
	parent.children[asByte-97] = NewTrieNode(word[0], parent)
	return RecurseInsert(word[1:], parent.children[asByte-97])
}

// func InsertNode(word string, parent *TrieNode) *TrieNode {
// 	asciiBytes := []byte(word)
// 	child := NewTrieNode(word[0], parent)
// 	parent.children[asciiBytes[0]-97] = child
// 	return child
// }