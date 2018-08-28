package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

// TODO: generate hash codes dynamically
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Generated code.</h1>")
	email := []byte("malinwolfhess@mac.com28-08-2018MalinStormHess")
	hash := sha256.Sum256(email)
	fmt.Fprintf(w, "%X ", hash)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)

}
