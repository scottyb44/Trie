package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// for reference
// a b c d e f g h i j k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25

func TestInsert_Dog(t *testing.T) {
	trie := NewTrie()

	trie.InsertWord("dog")

	require.Nil(t, trie.root.children[0])
	require.Equal(t, byte('d'), trie.root.children[3].value)
	require.Equal(t, byte('o'), trie.root.children[3].children[14].value)
	require.Equal(t, 1, 1)
}

func TestInsert_Book(t *testing.T) {
	trie := NewTrie()

	trie.InsertWord("book")

	require.Nil(t, trie.root.children[0])
	require.Equal(t, byte('b'), trie.root.children[1].value)
	require.Equal(t, byte('o'), trie.root.children[1].children[14].value)
	require.Equal(t, byte('o'), trie.root.children[1].children[14].children[14].value)
	require.Equal(t, 1, 1)
}

func TestFind_Dog(t *testing.T) {
	trie := NewTrie()
	trie.InsertWord("dog")

	found, node := trie.FindWord("dog")

	require.True(t, found)
	require.NotNil(t, node)
}

func TestFind_WordWithSimilarSpelling(t *testing.T) {
	trie := NewTrie()
	trie.InsertWord("book")
	trie.InsertWord("booker")

	found, node := trie.FindWord("book")

	require.True(t, found)
	require.NotNil(t, node)
}
