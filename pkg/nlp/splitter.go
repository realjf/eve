package nlp

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/fatih/set"
)

const SAME = 100
const VERY_LONG = 1000

const (
	SPLITTER_GENERAL = 1 + iota
	SPLITTER_MARKERS
	SPLITTER_SENT_END
	SPLITTER_SENT_START
)

type Splitter struct {
	SPLIT_AllowBetweenMarkers bool
	SPLIT_MaxWords            int64
	starters                  *set.Set
	enders                    map[string]bool
	markers                   map[string]int
}

func NewSplitter(splitterFile string) *Splitter {
	this := Splitter{
		starters: set.New(set.ThreadSafe).(*set.Set),
		enders:   make(map[string]bool),
		markers:  make(map[string]int),
	}

	cfg := NewConfigFile(false, "##")
	cfg.AddSection("General", SPLITTER_GENERAL)
	cfg.AddSection("Markers", SPLITTER_MARKERS)
	cfg.AddSection("SentenceEnd", SPLITTER_SENT_END)
	cfg.AddSection("SentenceStart", SPLITTER_SENT_START)

	if !cfg.Open(splitterFile) {
		CRASH("Error opening file"+splitterFile, MOD_SPLITTER)
	}

	this.SPLIT_AllowBetweenMarkers = true
	this.SPLIT_MaxWords = 0

	nmk := 1
	line := ""

	for cfg.GetContentLine(&line) {
		items := Split(line, " ")
		switch cfg.GetSection() {
		case SPLITTER_GENERAL:
			{
				name := items[0]
				if name == "AllowBetweenMarkers" {
					this.SPLIT_AllowBetweenMarkers, _ = strconv.ParseBool(items[1])
				} else if name == "MaxWords" {
					this.SPLIT_MaxWords, _ = strconv.ParseInt(items[1], 10, 64)
				} else {
					LOG.Panic("Unexpected splitter option " + name)
				}
				break
			}
		case SPLITTER_MARKERS:
			{
				open := items[0]
				close := items[1]
				if open != close {
					this.markers[open] = nmk
					this.markers[close] = -nmk
				} else {
					this.markers[open] = SAME + nmk
					this.markers[close] = SAME + nmk
				}
				nmk++
				break
			}
		case SPLITTER_SENT_END:
			{
				name := items[0]
				value, _ := strconv.ParseBool(items[1])
				this.enders[name] = !value
				break
			}
		case SPLITTER_SENT_START:
			{
				this.starters.Add(line)
				break
			}
		default:
			break
		}
	}

	LOG.Trace("Analyzer successfully created")
	return &this
}

type SplitterStatus struct {
	BetweenMark  bool
	NoSplitCount int
	MarkType     *list.List
	MarkForm     *list.List
	buffer       *Sentence
	nsentence    int
}

func (this *Splitter) OpenSession() *SplitterStatus {
	LOG.Trace("Opening new session")
	return &SplitterStatus{
		BetweenMark:  false,
		NoSplitCount: 0,
		MarkType:     list.New(),
		MarkForm:     list.New(),
		buffer:       NewSentence(),
		nsentence:    0,
	}
}

func (this *Splitter) CloseSession(ses *SplitterStatus) {
	LOG.Trace("Closing session")
	ses.MarkType = ses.MarkType.Init()
	ses.MarkForm = ses.MarkForm.Init()
	ses = nil
}

func (this *Splitter) Split(st *SplitterStatus, v *list.List, flush bool, ls *list.List) {
	// @todo
}

func (this *Splitter) endOfSentence(w *list.Element, v *list.List) bool {
	if w == v.Back() {
		return true
	} else {
		r := w
		r = r.Next()
		f := r.Value.(*Word).getForm()

		return strings.Title(f) == f || this.starters.Has(f)
	}
}
