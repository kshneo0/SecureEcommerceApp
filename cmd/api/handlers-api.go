package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

// GetPaymentIntent gets a payment intent, and returns json (or error json)
func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	//임시로 데이터 보내기
	j := jsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}
