package main

import (
	"fmt"
	"net/http"
)

// ServeDocs ...
func ServeDocs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
