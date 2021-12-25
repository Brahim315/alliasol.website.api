package handler

import (
	"fmt"
	"net/http"
)

func DevisHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "homepage")
}
