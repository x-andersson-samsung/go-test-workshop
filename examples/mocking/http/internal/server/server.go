package server

import (
	"io"
	"net/http"
	"strconv"
)

// Server implements a very simple server which returns the size of request body in bytes.
type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(len(reqBody))))
}
