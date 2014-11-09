package trie

import (
	"bytes"
	"unicode/utf8"
)

// The Node type makes up the Trie data structure.
type Node map[rune]Node

// New allocates a new node.
func New() Node { return make(Node) }

// Insert adds a new word to the Trie. It iterates over s and creates or appends
// a new Node for each rune.
func (node Node) Insert(s string) {
	for _, r := range s {
		if node[r] == nil {
			node[r] = New()
		}
		node = node[r]
	}
}

// Search looks for the string s inside the Trie structure. If s is matched and
// the node has mode children, its children are also returned as they all match
// the given prefix s.
func (node Node) Search(s string) []string {
	var res []string
	for i, r := range s {
		if _, ok := node[r]; ok {
			node = node[r]
		}
		if i == len(s)-utf8.RuneLen(r) {
			res = append(res, s)
		}
	}
	return append(res, node.All(s)...)
}

// String satisfies the Stringer interface for easily printing a Trie.
func (node Node) String() string {
	return node.print(0)
}

// All returns all the strings indexed by the current node and it's children in
// an array each item prefixed by p.
func (node Node) All(p string) (r []string) {
	for rn, c := range node {
		if len(c) > 0 {
			r = append(r, c.All(p+string(rn))...)
		}
	}
	return
}

func (node Node) print(pad int) string {
	if len(node) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	for rn, child := range node {
		for i := 0; i < pad; i++ {
			buf.WriteByte(' ')
		}
		buf.WriteRune(rn)
		buf.WriteByte('\n')
		if child != nil {
			buf.WriteString(child.print(pad + 1))
		}
	}
	return buf.String()
}
