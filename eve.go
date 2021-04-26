package main

import (
	. "github.com/realjf/eve/pkg/lib"
	. "github.com/realjf/eve/pkg/net"
	. "github.com/realjf/eve/terminal"
)

var logo = `            
	   ___           ___
	  / _ \ __  __  / _ \
	 |  __/\  \/  /|  __/ 
	  \___| \____/  \___|

	AdvancedLogic 2021 - v.0.1
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
