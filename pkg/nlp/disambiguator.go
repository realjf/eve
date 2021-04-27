package nlp

import (
	"container/list"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fatih/set"
)

const (
	DOCUMENT_SCOPE = iota
	SENTENCE_SCOPE
	ND_SCOPE
	SENTENCE_BIND
)

type Disambiguator struct {
	wnids map[string]*Synset
	binds map[string]*set.Set
}

func NewDisambiguator(disFile string) *Disambiguator {
	this := Disambiguator{
		wnids: make(map[string]*Synset),
		binds: make(map[string]*set.Set),
	}

	fileString, err := ioutil.ReadFile(disFile)
	if err != nil {
		LOG.Panic("Error loading file " + disFile)
	}
	lines := strings.Split(string(fileString), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		items := Split(line, "\t")
		sscope := items[0]
		scope := DOCUMENT_SCOPE
		if sscope == "d" {
			scope = DOCUMENT_SCOPE
		} else if sscope == "sd" {
			scope = SENTENCE_SCOPE
		} else if sscope == "sb" {
			scope = SENTENCE_BIND
		} else if sscope == "nd" {
			scope = ND_SCOPE
		}
		switch scope {
		case DOCUMENT_SCOPE, SENTENCE_SCOPE, ND_SCOPE:
			{
				lemma := items[1]
				wnid := items[2]
				pos, _ := strconv.ParseFloat(items[3], 64)
				neg, _ := strconv.ParseFloat(items[4], 64)
				domain := items[5][1:]
				score, _ := strconv.Atoi(items[6])
				gloss := items[7]
				syn := NewSynset(scope, lemma, wnid, pos, neg, domain, score, gloss)
				this.wnids[wnid] = syn
				break
			}
		case SENTENCE_BIND:
			{
				key := items[1][1:]
				for i := 2; i < len(items); i++ {
					if this.binds[key] == nil {
						this.binds[key] = set.New(set.ThreadSafe).(*set.Set)
					}
					this.binds[key].Add(items[i][1:])
				}
			}
		}
	}

	return &this
}

func (this *Disambiguator) Analyze(ss *list.List) {
	for s := ss.Front(); s != nil; s = s.Next() {
		for w := s.Value.(*Sentence).Front(); w != nil; w = w.Next() {
			lsens := w.Value.(*Word).getSenses(0)
			for l := lsens.Front(); l != nil; l = l.Next() {
				pair := l.Value.(FloatPair)
				id := pair.first
				prob := pair.second
				if prob > 0 {
					if this.wnids[id] != nil {
						println(id, w.Value.(*Word).getForm(), this.wnids[id].lemma, this.wnids[id].gloss, this.wnids[id].domain)
					}
				}
			}
		}
	}
}
