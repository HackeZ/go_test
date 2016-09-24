package main

import (
	r "go_test/gin_test/router"
)

// You Should visit
// http://127.0.0.1:8080/welcome/name/{yourname}/action/{whatyoudoing}?mess={andthemessage}

func main() {
	// r.RunRouterDefault(":8080")
	r.RunRouterMiddleware(":8080")
}
