/**
 * Example of simple http server deployed
 */
package main

import (
	"fmt"
	"net/http"
)

const LISTENING_PORT string = "8080"

/**
 * main function
 */
func main() {
	fmt.Printf("Starting server on port %s...", LISTENING_PORT)
	http.HandleFunc("/", handler)                // Bind handler on pattern "/"
	http.ListenAndServe(":"+LISTENING_PORT, nil) // Handler is nil => Default use DefaultServeMux
	fmt.Printf("Shutting down server.")
}

/**
 * handler function
 */
func handler(w http.ResponseWriter, r *http.Request) {
	infoToRender := "1892"
	fmt.Fprintf(w, "Hi there, here is your information : %s", infoToRender)
}
