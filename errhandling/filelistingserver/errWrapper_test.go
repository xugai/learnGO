package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"./filelisting"
	"os"
	"errors"
	"fmt"
)

var tests = []struct {
	h appHandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, http.StatusInternalServerError, "User Error"},
	{errNotFound, http.StatusNotFound, "Not Found"},
	{errNotPermission, http.StatusForbidden, "Forbidden"},
	{errUnknownErr, 500, "Internal Server Error"},
	{normalCase, 200, ""},
}


func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return filelisting.UserError("User Error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknownErr(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("Unknown error")
}

func normalCase(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(writer, "no error")
	return nil
}


func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	body, _ := ioutil.ReadAll(resp.Body)
	b := strings.Trim(string(body), "\n")
	if resp.StatusCode != expectedCode || b != expectedMsg {
		t.Errorf("except (%d, %s), actual (%d, %s)",
			expectedCode, expectedMsg, resp.StatusCode, b)
	}
}
