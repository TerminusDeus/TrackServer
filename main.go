// TrackServer project main.go
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	encoded := r.URL.Path[1:]
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Fprintf(w, "decoded string: %s \ndecode error: ", string(decoded), err)
		return
	}
	http.Redirect(w, r, string(decoded), http.StatusFound)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8085", nil)
}
