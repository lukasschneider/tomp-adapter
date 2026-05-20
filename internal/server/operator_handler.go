package server

import "net/http"

// Describes the status of the Transport Operator - whether the APIs are running or not
func (s *Server) OperatorPingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// describes the running implementations
func (s *Server) OperatorMetaHandler(w http.ResponseWriter, r *http.Request) {
}
