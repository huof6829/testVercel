package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
	// fmt.Fprintf(w, "r.url:%+v, host:%+v\n", r.URL.Path, r.Host)
	// fmt.Fprintf(w, "r:%+v\n", r)
	fmt.Fprintf(w, "w:%+v\n", w)
	fmt.Fprintf(w, "w.header:%+v\n", w.Header())

}
