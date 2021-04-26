package main

import (
	. "github.com/realjf/eve/pkg/lib"
	. "github.com/realjf/eve/pkg/net"
	. "github.com/realjf/eve/terminal"
)

var logo = `            
	  ___  __    __  ___
	 / _ \\  \  /  // _ \
	|  __/ \  \/  /|  __/ 
	 \___|  \____/  \___|

	realjf.io 2021 - v0.0.1
`

func init() {
	Infof("eve - Natural Language Processing for Golang\n")
}

func main() {
	analyzer := NewAnalyzer()

	println(logo)

	httpServer := NewHttpServer(analyzer)
	httpServer.Listen()
}
