package main

import (
	_ "./mis"
	"net/http"
)

func main() {
	http.ListenAndServe(":8007", nil)
}
