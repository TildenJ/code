package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)

	_ = os.Setenv("HTTP_PROXY", "http://127.0.0.1:8910")
	_ = os.Setenv("HTTPS_PROXY", "http://127.0.0.1:8910")
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
