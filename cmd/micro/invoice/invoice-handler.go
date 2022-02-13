package main

import (
	"fmt"
	"net/http"
)

// CreateAndSendInvoice creates an invoice as a PDF, and emails it to recipient
func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", "world")
}
