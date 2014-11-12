package trie

import (
	"reflect"
	"sort"
	"testing"
	"unicode/utf8"
)

func TestInsert(t *testing.T) {
	trie := New()
	trie.Insert("ab")
	trie.Insert("ac")
	expected := Node{'a': Node{'b': Node{}, 'c': Node{}}}
	expected['a']['b'].End()
	expected['a']['c'].End()
	if !reflect.DeepEqual(trie, expected) {
		t.Errorf("unequal trie %q, %q", trie, expected)
	}
}

func TestAll(t *testing.T) {
	for _, test := range []struct {
		words []string
	}{
		{[]string{"ab", "ac"}},
		{[]string{"a", "ab", "ac"}},
		{[]string{"a", "ab", "ac", "abc"}},
		{[]string{"a", "ab", "ac", "abc", "abcdef"}},
	} {
		trie := New()
		for _, word := range test.words {
			trie.Insert(word)
		}
		all := trie.All("")
		sort.Strings(all)
		sort.Strings(test.words)
		if !reflect.DeepEqual(all, test.words) {
			t.Errorf("when calling all on trie the result set should be %q not %q", test.words, all)
		}
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
		sort.Strings(results)
		sort.Strings(test.results)
		if !reflect.DeepEqual(test.results, results) {
			t.Errorf("when searching for %q the result set should be %q not %q", test.search, test.results, results)
		}
	}
}

func TestEnd(t *testing.T) {
	if utf8.ValidRune(end) {
		t.Errorf("end variable should not use a valid rune")
	}
}
