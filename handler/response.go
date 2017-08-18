package handler

import (
	"encoding/json"
	"net/http"
)

// postResult is JSON response returned as reply on POST and in case of errors
type postResult struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

// returnResult Sends postResult
func returnResult(writer http.ResponseWriter, errStr string) {
	var res postResult

	if errStr != "" {
		res.Error = errStr
	} else {
		res.OK = true
	}

	payload, err := json.Marshal(res)

	if err == nil {
		writer.Write(payload)
	}
}
