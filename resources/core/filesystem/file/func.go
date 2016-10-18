package file

import (
	"cici/api"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func get(response http.ResponseWriter, request *http.Request) {
	var err error

	path := request.URL.Query().Get("path")

	if path == "" {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: 400,
			Message:    "path query param must be supplied",
		})
		return
	}

	encoding := request.URL.Query().Get("encoding")

	if encoding != "base32" && encoding != "base64" && encoding != "hex" && encoding != "" {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: 400,
			Message:    "encoding query param must be of: base32, base64 or hex",
		})
		return
	}

	download := request.URL.Query().Get("download")

	file, err := os.Stat(path)

	if err != nil {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: 400,
			Message:    err.Error(),
		})
		return
	}

	if file.IsDir() {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: 400,
			Message:    "Cannot read contents of a directory as a file",
		})
		return
	}

	filebytes, err := ioutil.ReadFile(path)

	if err != nil {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: 500,
			Message:    err.Error(),
		})
		return
	}

	var contentType = request.Header.Get("Accept")

	filecontent := &fileContent{
		Path:     path,
		Length:   file.Size(),
		Encoding: encoding,
		Content:  encode(filebytes, encoding),
	}

	if download == "true" {
		response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%v"`, file.Name()))
	}

	isHTMLRequested := strings.Contains(contentType, "text/html")
	isXMLRequested := strings.Contains(contentType, "application/xml")
	isJSONRequested := strings.Contains(contentType, "application/json")

	if isHTMLRequested {
		response.Header().Set("Content-Type", "text/html")
		api.WriteResponse(response, 200, []byte(filecontent.Content))
	} else if isXMLRequested {
		b, _ := xml.Marshal(filecontent)
		response.Header().Set("Content-Type", "application/xml")
		api.WriteResponse(response, 200, b)
	} else if isJSONRequested {
		b, _ := json.Marshal(filecontent)
		response.Header().Set("Content-Type", "application/json")
		api.WriteResponse(response, 200, b)
	} else {
		response.Header().Set("Content-Type", "text/plain")
		api.WriteResponse(response, 200, []byte(filecontent.Content))
	}
}

func encode(b []byte, e string) string {
	switch e {
	case "base32":
		return base32.StdEncoding.EncodeToString(b)
	case "base64":
		return base64.StdEncoding.EncodeToString(b)
	case "hex":
		return hex.EncodeToString(b)
	default:
		return string(b)
	}
}
