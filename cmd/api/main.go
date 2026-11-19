package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Veer0729/olx_golang/internal/config"
)

func main(){

	cfg:= config.MustLoad()


	fmt.Println("server is starting...")
	// 1. Create router (mux)
	mux:= http.NewServeMux() // Our private router

	// 2. Register routes (HandleFunc)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request){
		// r tells you what the client wants, w is how you give them the answer.

		w.Header().Set("Content-Type", "application/json") 
		// w.Header() gives you a map of response headers. You're adding one header that tells the client: "Don't try to parse this as HTML or plain text — it's JSON."

		w.WriteHeader(http.StatusOK)
		// http.StatusOK is just a named constant for the number 200. You're explicitly saying "this request succeeded."

		w.Write([]byte(`{"status": "ok"}`))
	})
	  
	// 3. Configure server (http.Server{...})
	srv:= http.Server{
		Addr: ":" + cfg.Port, // "Which port do I listen on?"
		Handler: mux, // "Who do I hand requests to?"
		ReadTimeout: time.Second * 10, // "How long do I wait for the client to finish sending their request (headers + body)?"
		WriteTimeout: time.Second * 30, // "How long do I give my handler to finish writing the response?" 
		IdleTimeout: time.Second * 60, // "How long do I keep an idle keep-alive connection open before closing it?"
	}

	// 4. Start it and die if it can't (ListenAndServe)
	if err:= srv.ListenAndServe(); err != nil{ // ListenAndServe: Starts the server — it begins listening on :8090 and blocks (never returns) while the server is running
		log.Fatalf("server failed: %v", err)
	}
}