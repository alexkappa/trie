package trie

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	trie := New()
	trie.Insert("ab")
	trie.Insert("ac")
	expected := Node{'a': Node{'b': Node{}, 'c': Node{}}}
	if !reflect.DeepEqual(trie, expected) {
		t.Errorf("unequal trie %q, %q", trie, expected)
	}
}

func TestSearch(t *testing.T) {
	for _, test := range []struct {
		dict    []string
		search  string
		results []string
	}{
		{[]string{"ab", "ac"}, "ab", []string{"ab"}},
		{[]string{"a", "ab", "ac"}, "a", []string{"a", "ab", "ac"}},
		{[]string{"a", "ab", "ac", "abc"}, "ab", []string{"ab", "abc"}},
		{[]string{"a", "ab", "ac", "abc", "abcdef"}, "ab", []string{"ab", "abc", "abcdef"}},
	} {
		trie := New()
		for _, word := range test.dict {
			trie.Insert(word)
		}
		results := trie.Search(test.search)
		if !reflect.DeepEqual(test.results, results) {
			t.Errorf("when searching for %q the result set should be %q not %q", test.search, test.results, results)
		}
	}
}
