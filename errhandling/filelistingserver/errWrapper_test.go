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
)

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
	return nil
}


func TestErrWrapper(t *testing.T) {
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

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		body, _ := ioutil.ReadAll(response.Body)
		b := strings.Trim(string(body), "\n")
		if response.Code != tt.code || b != tt.message {
			t.Errorf("except (%d, %s), actual (%d, %s)",
						tt.code, tt.message, response.Code, b)
		}
	}
}
