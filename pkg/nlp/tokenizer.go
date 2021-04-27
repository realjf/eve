package nlp

import (
	"container/list"

	"github.com/fatih/set"
)

const (
	TOKENIZER_MACROS = 1 + iota
	TOKENIZER_REGEXPS
	TOKENIZER_ABBREV
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
	cfg.AddSection("Macros", TOKENIZER_MACROS)
	cfg.AddSection("RegExps", TOKENIZER_REGEXPS)
	cfg.AddSection("Abbreviations", TOKENIZER_ABBREV)

	if !cfg.Open(tokenizerFile) {
		LOG.Panic("Error opening file " + tokenizerFile)
	}

	macros := list.New()
	rul := false
	var ci string
	line := ""
	for cfg.GetContentLine(&line) {
		items := Split(line, " ")
		switch cfg.GetSection() {
		case TOKENIZER_MACROS:
			{
				if rul {
					LOG.Panic("Error reading tokenizer configuration. Macros must be defined before rules.")
				}
				mname := items[0]
				mvalue := items[1]
				macros.PushBack(Pair{mname, mvalue})
				LOG.Trace("Read macro " + mname + ": " + mvalue)
				break
			}
		case TOKENIZER_REGEXPS:
			{
				var substr int
				comm := items[0]

			}
		}
	}

	return &this
}
