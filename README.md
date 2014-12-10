# Trie

Trie implementation in Go. Inspired by John Resig's [trie-js](https://github.com/jeresig/trie-js).

## Motivation

The trie data structure is particularly interresting to me as it's surprisingly simple yet powerful.

The data structure is nothing more than the recursive type `type Node map[rune]Node`. With this simple type we're able to index words by prefix and perform fast lookups.

## Usage

```Go
import "github.com/alexkappa/trie"

t := trie.New()
t.Index([]string{"ab", "ac", "ad", "abc"})
t.Search("ab") // ["ab" "abc"]
```

The API documentation is available at [godoc.org](http://godoc.org/github.com/alexkappa/trie).