package common

import (
    "net/http"
    "strconv"
    "encoding/json"
)

const kInternalError = "An internal error occurred"

type ErrorResponse struct {
    Error string `json:"error"`
}


func WriteJsonResponse(w http.ResponseWriter, status int, data []byte) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Content-Length", strconv.Itoa(len(data)))
    w.WriteHeader(status)
    w.Write(data)
}


func WriteInternalErrorResponse(w http.ResponseWriter) {
    data, _ := json.Marshal(ErrorResponse{kInternalError})
    WriteJsonResponse(w, http.StatusInternalServerError, data)
}
