package nlp

import (
	"github.com/kdar/factorlog"
)

func init() {

}

type NLPOptions struct {
	Serverity         factorlog.Serverity
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

func NewNLPOptions() *NLPOptions {
	return &NLPOptions{}
}
