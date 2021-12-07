package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	FailedToParseAccountID               = "Couldn't parse the accountId, make sure it's an integer!"
	FailedToMarshalResponsePayloadToJson = "Couldn't marshal the payload to jsonPayload!"
)

type Server struct {
	httpServer *http.Server
	port       string
}

type responsePayload struct {
	AccountID int64     `json:"accountId"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

func new(port string) *Server {
	return &Server{
		port:       port,
		httpServer: &http.Server{Addr: port},
	}
}

func (s *Server) run() error {
	s.httpServer.Handler = s.handler()
	return s.httpServer.ListenAndServe()
}

func (s *Server) handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.accountId)
	mux.HandleFunc("/healtz", s.healthCheck)
	return mux
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) accountId(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.RequestURI(), "/")
	accountId := tmp[1]
	fmt.Printf("accountId: %s\n", accountId)

	// convert string to int
	accountIdInt, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil {
		fmt.Println(err)
		http.Error(w, FailedToParseAccountID, http.StatusBadRequest)
		return
	}

	// create response payload
	payload := &responsePayload{
		AccountID: accountIdInt,
		Timestamp: time.Now().UTC(),
		Data:      RandomText,
	}

	// marshal payload to jsonPayload
	jsonPayload, err := json.Marshal(payload)
	// check for errors
	if err != nil {
		fmt.Println(err)
		http.Error(w, FailedToMarshalResponsePayloadToJson, http.StatusInternalServerError)
		return
	}

	// write jsonPayload to response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPayload)
}
