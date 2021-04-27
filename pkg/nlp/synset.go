package nlp

import "strings"

type Synset struct {
	scope    int
	lemma    string
	wnid     string
	shortTag string
	pos      float64
	neg      float64
	domain   string
	score    int
	gloss    string
}

func NewSynset(scope int, lemma string, wnid string, pos float64, neg float64, domain string, score int, gloss string) *Synset {
	shortTag := wnid[strings.Index(wnid, "-")+1:]
	return &Synset{
		scope:    scope,
		lemma:    lemma,
		wnid:     wnid,
		shortTag: shortTag,
		pos:      pos,
		neg:      neg,
		domain:   domain,
		score:    score,
		gloss:    gloss,
	}
}
