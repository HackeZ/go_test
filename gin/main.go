package main

import (
	r "go_test/gin/router"
)

// You Should visit
// http://127.0.0.1:8080/welcome/name/{yourname}/action/{whatyoudoing}?mess={andthemessage}

const (
	// ServerAddrPort Default Port
	ServerAddrPort = ":8080"
)

func main() {
	// r.RunRouterDefault(ServerAddrPort)
	r.RunRouterMiddleware(ServerAddrPort)
	// r.RunRouterWithModel(ServerAddrPort)
	// r.RunRouterRendering(ServerAddrPort)
}
