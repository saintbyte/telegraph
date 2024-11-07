package api

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func Env(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Fprintf(w, pair[0])
		fmt.Fprintf(w, "=")
		fmt.Fprintf(w, pair[1])
		fmt.Fprintf(w, "<br />")
	}
	fmt.Fprintf(w, "GOMAXPROCS")
	fmt.Fprintf(w, "=")
	fmt.Fprintf(w, "%v", runtime.NumCPU())
	fmt.Fprintf(w, "<br />")
}
