package api

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

// WriteErrorResult writes a error result as an http response
func WriteErrorResult(request *http.Request, response http.ResponseWriter, result *ErrorResult) {
	var b []byte

	contentType := request.Header.Get("Accept")

	switch {
	case strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/json"):
		b, _ = json.Marshal(result)
		response.Header().Set("Content-Type", contentType)

	case strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml"):
		b, _ = xml.Marshal(result)
		response.Header().Set("Content-Type", contentType)

	default:
		b = []byte(result.Message)
		response.Header().Set("Content-Type", "text/plain")
	}

	response.WriteHeader(result.StatusCode)
	respi, err := response.Write(b)

	if err != nil {
		log.Println("Error writing http response:", respi, err)
	}
}

// WriteSuccessResult ...
func WriteSuccessResult(request *http.Request, response http.ResponseWriter, result *Result) {
	var b []byte

	contentType := request.Header.Get("Accept")

	switch {
	case contentType == "application/json" || contentType == "text/json":
		b, _ = json.Marshal(result)
		response.Header().Set("Content-Type", contentType)

	case contentType == "application/xml" || contentType == "text/xml":
		b, _ = xml.Marshal(result)
		response.Header().Set("Content-Type", contentType)

	default:
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		enc.Encode(result.Data)
		b = buf.Bytes()
		response.Header().Set("Content-Type", "text/html")
	}

	response.WriteHeader(result.StatusCode)
	respi, err := response.Write(b)

	if err != nil {
		log.Println("Error writing http response:", respi, err)
	}
}

// WriteResponse ...
func WriteResponse(response http.ResponseWriter, status int, content []byte) {
	response.WriteHeader(status)
	response.Write(content)
}
