package lib

import (
	. "github.com/realjf/eve/pkg/engine"
)

type Analyzer struct {
	context *Context
}

func NewAnalyzer() *Analyzer {
	context := NewContext("conf/eve.toml")
	// context.InitNLP()
	instance := new(Analyzer)
	instance.context = context

	return instance
}

func (a *Analyzer) Int64(key string, def int64) int64 {
	return a.context.Int64(key, def)
}

func (a *Analyzer) AnalyzeText() {

}
