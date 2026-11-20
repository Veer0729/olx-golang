package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request){
	// r tells you what the client wants, w is how you give them the answer.

	w.Header().Set("Content-Type", "application/json") 
	// w.Header() gives you a map of response headers. You're adding one header that tells the client: "Don't try to parse this as HTML or plain text — it's JSON."

	w.WriteHeader(http.StatusOK)
	// http.StatusOK is just a named constant for the number 200. You're explicitly saying "this request succeeded."

	w.Write([]byte(`{"status": "ok"}`))
}