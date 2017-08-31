package trie

import (
	"bytes"
	"sort"
	"unicode/utf8"
)

const end = -1

// The Node type makes up the Trie data structure.
type Node map[rune]Node

// New allocates a new node.
func New() Node { return make(Node) }

// Index builds a Trie with the supplied dictionary d.
func (node Node) Index(d []string) {
	for _, s := range d {
		node.Insert(s)
	}
}

// Insert adds a new word to the Trie. It iterates over s and creates or appends
// a new Node for each rune.
func (node Node) Insert(s string) {
	for _, r := range s {
		if node[r] == nil {
			node[r] = New()
		}
		node = node[r]
	}
	node.end()
}

// Search looks for the string s inside the Trie structure. If s is matched and
// the node has mode children, its children are also returned as they all match
// the given prefix s.
func (node Node) Search(s string) (res []string) {
	for i, r := range s {
		if _, ok := node[r]; ok {
			node = node[r]
		}
		if i == len(s)-utf8.RuneLen(r) {
			return append(res, node.All(s)...)
		}
	}
	return res
}

// All returns all the strings indexed by the current node and it's children in
// an array each item prefixed by p.
func (node Node) All(p string) (res []string) {
	if node.IsEnd() {
		res = append(res, p)
	}
	node.ForEach(func(r rune, n Node) {
		if len(n) > 0 {
			res = append(res, n.All(p+string(r))...)
		}
	})
	return
}

// End marks the current node as the end of an indexed word. For example when
// indexing the string "abc", the node pointed by the key 'c' is marked as an
// end node.
func (node Node) end() {
	node[end] = nil
}

// IsEnd returns true if the current node is an end node.
func (node Node) IsEnd() bool {
	_, found := node[end]
	return found
}

// Iter is a function type used as an argument to ForEach. It's arguments are
// the same as what would be returned on each for loop cycle, namely a rune and
// a map of nodes.
type Iter func(r rune, n Node)

// ForEach wraps the for loop and additionally checks for the end rune and
// ignores it.
func (node Node) ForEach(f Iter) {

	keys := make([]rune, 0, len(node))
	for key, _ := range node {
		keys = append(keys, key)
	}

	sort.Sort(runeSlice(keys))

	for _, key := range keys {
		if key == end {
			continue
		}
		f(key, node[key])
	}
}

// String satisfies the Stringer interface for easily printing a Trie.
func (node Node) String() string {
	return node.print(0)
}

// Prints each item in the node prefixed with pad spaces. For each nested node
// print is called recursively with an incremented pad.
func (node Node) print(pad int) string {
	if len(node) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	node.ForEach(func(r rune, n Node) {
		for i := 0; i < pad; i++ {
			buf.WriteByte(' ')
		}
		buf.WriteRune(r)
		buf.WriteByte('\n')
		if n != nil {
			buf.WriteString(n.print(pad + 1))
		}
	})
	return buf.String()
}

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Less(i, j int) bool { return r[i] < r[j] }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
