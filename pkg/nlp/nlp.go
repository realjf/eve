package nlp

import (
	"container/list"
	"strings"

	"github.com/fatih/set"
	"github.com/kdar/factorlog"

	"github.com/realjf/eve/pkg/models"
	"github.com/realjf/eve/pkg/wordnet"
)

var LOG *factorlog.FactorLog

const (
	PANIC   = factorlog.PANIC
	FATAL   = factorlog.FATAL
	ERROR   = factorlog.ERROR
	WARN    = factorlog.WARN
	DEBUG   = factorlog.DEBUG
	INFO    = factorlog.INFO
	VERBOSE = factorlog.TRACE

	TAG_NP = "NP"
)

func init() {

}

type NLPOptions struct {
	Serverity         factorlog.Severity
	DataPath          string
	Lang              string
	TokenizerFile     string
	SplitterFile      string
	MorfoOptions      *MacoOptions
	TaggerFile        string
	ShallowParserFile string
	SenseFile         string
	UKBFile           string
	DisambiguatorFile string
	Status            func()
}

func NewNLPOptions(dataPath string, lang string, f func()) *NLPOptions {
	return &NLPOptions{
		DataPath: dataPath,
		Lang:     lang,
		Status:   f,
	}
}

type NLPEngine struct {
	options       *NLPOptions
	tokenizer     *Tokenizer
	splitter      *Splitter
	morfo         *Maco
	tagger        *HMMTagger
	grammar       *Grammar
	shallowParser *ChartParser
	sense         *Senses
	dsb           *UKB
	disambiguator *Disambiguator
	filter        *set.Set
	mitie         *MITIE
	WordNet       *wordnet.WN
}

func NewNLPEngine(options *NLPOptions) *NLPEngine {
	this := NLPEngine{
		options: options,
	}

	LOG.SetMinMaxSeverity(factorlog.PANIC, options.Serverity)

	if options.TokenizerFile != "" {
		this.tokenizer = NewTokenizer(options.DataPath + "/" + options.Lang + "/" + options.TokenizerFile)
		this.options.Status()
	}

	if options.SplitterFile != "" {
		this.splitter = NewSplitter(options.DataPath + "/" + options.Lang + "/" + options.SplitterFile)
		this.options.Status()
	}

	if options.MorfoOptions != nil {
		this.morfo = NewMaco(options.MorfoOptions)
		this.options.Status()
	}

	if options.SenseFile != "" {
		this.sense = NewSenses(options.DataPath + "/" + options.Lang + "/" + options.SenseFile)
		this.options.Status()
	}

	if options.TaggerFile != "" {
		this.tagger = NewHMMTagger(options.DataPath+"/"+options.Lang+"/"+options.TaggerFile, true, FORCE_TAGGER, 1)
		this.options.Status()
	}

	if options.ShallowParserFile != "" {
		this.grammar = NewGrammar(options.DataPath + "/" + options.Lang + "/" + options.ShallowParserFile)
		this.shallowParser = NewChartParser(this.grammar)
		this.options.Status()
	}

	if options.UKBFile != "" {
		this.dsb = NewUKB(options.DataPath + "/" + options.Lang + "/" + options.UKBFile)
		this.options.Status()
	}

	if options.DisambiguatorFile != "" {
		if strings.HasPrefix(options.DisambiguatorFile, "common") {
			this.disambiguator = NewDisambiguator(options.DataPath + "/" + options.DisambiguatorFile)
		} else {
			this.disambiguator = NewDisambiguator(options.DataPath + "/" + options.Lang + "/" + options.DisambiguatorFile)
		}
		this.options.Status()
	}

	this.mitie = NewMITIE(options.DataPath + "/" + options.Lang + "/mitie/ner_model.dat")
	this.options.Status()

	return &this
}

func (this *NLPEngine) Workflow(document *models.DocumentEntity, output chan *models.DocumentEntity) {
	defer func() {
		if r := recover(); r != nil {
			err, _ := r.(error)
			if err != nil {
				output <- nil
			} else {
				output <- nil
			}
		}
	}()

	document.Init()
	tokens := list.New()
	url := document.Url
	content := document.Content

	if url != "" && content == "" {

	}

	// @todo

	output <- document
}

func (this *NLPEngine) PrintList(document *models.DocumentEntity) {
	ls := document.Sentences()
	for l := ls.Front(); l != nil; l = l.Next() {
		for w := l.Value.(*Sentence).Front(); w != nil; w = w.Next() {
			item := w.Value.(*Word).getForm() + ":"
			for a := w.Value.(*Word).Front(); a != nil; a = a.Next() {
				if a.Value.(*Analysis).isSelected(0) {
					item += a.Value.(*Analysis).getTag()
				}
			}
			println(item)
		}
	}
}

func (this *NLPEngine) PrintTree(document *models.DocumentEntity) {
	ls := document.Sentences()
	for l := ls.Front(); l != nil; l = l.Next() {
		tr := l.Value.(*models.SentenceEntity).GetSentence().(*Sentence).pts[0]
		output := new(Output)
		out := ""

		output.PrintTree(&out, tr.begin(), 0)

		LOG.Trace(out)
		println(out)
	}
}
