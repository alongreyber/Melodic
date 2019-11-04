package main
// Copied from
// https://medium.com/ki-labs-engineering/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"
)
// ClientError is an error whose details to be shared with client.
type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// HTTPError implements ClientError interface.
type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	response := make(map[string]interface{})
	response["error"] = e
	response["ok"] = false
	body, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, detail string, status int) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

func errorResponse(w http.ResponseWriter, err error) {
    	log.Printf("An error accured: %v", err) // Log the error.

	clientError, ok := err.(ClientError) // Check if it is a ClientError.
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}
