package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/cheggaaa/pb"
	"github.com/realjf/eve/pkg/nlp"
	. "github.com/realjf/eve/terminal"
)

type Engine struct {
	semaphore *sync.Mutex
	NLP       *nlp.NLPEngine
	Ready     bool
}

func NewEngine() *Engine {
	return &Engine{
		semaphore: new(sync.Mutex),
		Ready:     false,
	}
}

var (
	path = "./"
	lang = "en"
)

func (e *Engine) InitNLP() {
	e.semaphore.Lock()
	defer e.semaphore.Unlock()

	if e.Ready {
		return
	}

	Infoln("Init Natural Language Processing Engine")
	initialized := false
	count := 80 //
	bar := pb.StartNew(count)
	bar.ShowPercent = true
	bar.ShowCounters = false

	inc := func() {
		for i := 0; i < 10; i++ {
			bar.Increment()
		}
	}

	start := time.Now().UnixNano()
	//
	nlpOptions := nlp.NewNLPOptions(path+"data/", lang, inc)
	nlpOptions.Serverity = nlp.ERROR

	nlpEngine := nlp.NewNLPEngine(nlpOptions)

	stop := time.Now().UnixNano()
	delta := (stop - start) / (1000 * 1000)
	initialized = true

	bar.FinishPrint(fmt.Sprintf("Data loaded in %dms", delta))

	e.NLP = nlpEngine
	e.Ready = initialized
}
