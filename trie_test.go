package trie

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"unicode/utf8"
)

func TestIndex(t *testing.T) {
	trie := New()
	trie.Index([]string{"ab", "ac", "ad", "abc"})
	matches := trie.Search("ab")
	expected := []string{"ab", "abc"}
	sort.Strings(matches)
	sort.Strings(expected)
	if !reflect.DeepEqual(matches, expected) {
		t.Errorf("expected search results %v expected %v", matches, expected)
	}
}

func TestInsert(t *testing.T) {
	trie := New()
	trie.Insert("ab")
	trie.Insert("ac")
	expected := Node{'a': Node{'b': Node{}, 'c': Node{}}}
	expected['a']['b'].end()
	expected['a']['c'].end()
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

func TestString(t *testing.T) {
	trie := New()
	trie.Index([]string{"a", "ab", "abc"})
	expected := "a\n b\n  c\n"
	if trie.String() != expected {
		t.Errorf("unexpected string representation of trie %q expected %q", trie.String(), expected)
	}
}

func ExampleNode() {
	trie := New()
	trie.Index([]string{"ab", "ac", "ad", "abc"})
	fmt.Printf("%q", trie.Search("ab"))
	// Output: ["ab" "abc"]
}

func ExamplePrint() {
	trie := New()
	trie.Index([]string{"ab", "ac", "ad", "abc", "abcd"})
	fmt.Printf("%s", trie)
	// Output: a
	//  b
	//   c
	//    d
	//  c
	//  d
}
