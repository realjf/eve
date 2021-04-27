package nlp

import (
	"container/list"

	"github.com/fatih/set"
)

type Tokenizer struct {
	abrevs  *set.Set
	rules   *list.List
	matches map[string]int
}

func NewTokenizer(tokenizerFile string) *Tokenizer {
	this := Tokenizer{
		abrevs:  set.New(set.ThreadSafe).(*set.Set),
		rules:   list.New(),
		matches: make(map[string]int),
	}

	cfg := NewConfigFile(false, "##")

	return &this
}
