package main

import (
	"net/http"
	"fmt"
	"os"
)

// Run exec to start http server.
func serve() {
	port := 8080
	fmt.Fprintf(os.Stdout, "start to run http server at port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start server: %s", err)
		os.Exit(1)
	}
}

func main() {
	serve()
}