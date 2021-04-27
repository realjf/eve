package nlp

import (
	"container/list"
	"regexp"
	"strconv"
	"strings"

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
				substr, _ = strconv.Atoi(items[1])
				re := items[2]
				rul = true

				for i := macros.Front(); i != nil; i = i.Next() {
					mname := "{" + i.Value.(Pair).first.(string) + "}"
					mvalue := i.Value.(Pair).second.(string)
					p := strings.Index(re, mname)
					for p > -1 {
						re = strings.Replace(re, mname, mvalue, -1)
						p = strings.Index(re[p:], mname)
					}
				}

				if len(items) > 3 {
					ci = items[3]
				}

				if ci == "CI" {
					newre := "(?i)" + re
					x, err := regexp.Compile(newre)
					if err == nil {
						this.rules.PushBack(Pair{comm, x})
					} else {
						LOG.Warn("Rule " + comm + " [" + newre + "] failed to be compiled")
					}
				} else {
					x, err := regexp.Compile(re)
					if err == nil {
						this.rules.PushBack(Pair{comm, x})
					} else {
						LOG.Warn("Rule " + comm + " [" + re + "] failed to be compiled")
					}
				}

				this.matches[comm] = substr
				LOG.Trace("Stored rule " + comm + " " + re + " " + strconv.Itoa(substr))
				break
			}
		case TOKENIZER_ABBREV:
			{
				this.abrevs.Add(line)
				break
			}
		default:
			break
		}
	}

	LOG.Trace("analyzer successfully created")
	return &this
}

func (this *Tokenizer) Tokenize(p string, offset int, v *list.List) {
	// @todo
}
