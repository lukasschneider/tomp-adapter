package tomperr

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	ErrorCode int32  `json:"errorcode"`
	Type      string `json:"type,omitempty"`
	Title     string `json:"title,omitempty"`
	Status    int    `json:"status,omitempty"`
	Detail    string `json:"detail,omitempty"`
	Instance  string `json:"instance,omitempty"`
}

/*
x000 to x199 are for technical errors that shouldn't happen in the normal usage of the TOMP API.
x200 to x499 are for functional errors that are to be expected in the normal usage of the TOMP API.
x500 and onwards are not yet defined.

If code > 999 or < 0 is provided the function returns 0

if module > 9 or < 0 is provides the function return 0
*/
func NewErrorCode(module TompModule, code int) int32 {

	if module > 9 || module < 0 {
		return 0
	}

	if code > 999 || code < 0 {
		return 0
	}

	return int32(module)*1000 + int32(code)
}

// Returns a TOMP valid Error based on
// https://github.com/TOMP-WG/TOMP-API/wiki/Error-handling-in-TOMP
func httpError(w http.ResponseWriter, status int, e Error) {
	payload, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
