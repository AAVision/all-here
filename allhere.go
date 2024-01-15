package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func starter(userInput UserInput) {
	/**
	This is the starter of the server.
	@var UserInput UserInput.
	@return nil.
	*/
	port := userInput.Port
	if port == "" {
		port = "9876"
	} else {
		port = string(port)
	}

	fmt.Print("Server started at port: ")
	color.Yellow.Print(port)
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/{status_code}", setStatusCode)

	go http.ListenAndServe(":"+port, r)
	select {}
}

func setStatusCode(w http.ResponseWriter, r *http.Request) {
	/**
	This function to set Status Code.
	@var w http.ResponseWriter.
	@var r *http.Request.
	@return nil.
	*/
	vars := mux.Vars(r)
	urlStatusCode := vars["status_code"]

	statusCode, err := stringToInt(urlStatusCode)

	if err != nil {
		fmt.Println("Please provide a valid status code!")
		return
	}

	isValidRange := isValidStatusCodeRange(statusCode)

	if !isValidRange {
		fmt.Println("Please a status code with range between 100 and 599!")
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	statusText := http.StatusText(statusCode)

	resp := make(map[string]string)
	resp["message"] = statusText
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func loggingMiddleware(next http.Handler) http.Handler {
	/**
	This function will display server logger.
	@var next http.Handler.
	@return http.Handler.
	*/
	return handlers.LoggingHandler(os.Stdout, next)
}
