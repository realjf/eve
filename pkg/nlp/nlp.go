package nlp

import (
	"github.com/kdar/factorlog"
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
	options   *NLPOptions
	tokenizer *Tokenizer
}

func NewNLPEngine(options *NLPOptions) *NLPEngine {
	this := NLPEngine{
		options: options,
	}

	LOG.SetMinMaxSeverity(factorlog.PANIC, options.Serverity)

	return &this
}
