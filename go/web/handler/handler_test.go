/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午6:46
 * Description:
 **/

package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errCustomError, 400, "url path must be start with " + prefix},
	{errCustomPathError, 400, "url path illegal  "},
	{errNotFund, 404, "Not Found"},
	{errPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, ""},
}

func errPanic(w http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errCustomError(w http.ResponseWriter, request *http.Request) error {
	return userError("url path must be start with " + prefix)
}

func errCustomPathError(w http.ResponseWriter, request *http.Request) error {
	return userError("url path illegal  ")
}

func errNotFund(w http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errPermission(w http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(w http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(w http.ResponseWriter, request *http.Request) error {
	return nil
}

func TestAppHandler(t *testing.T) {

	for _, tt := range tests {
		f := AppHandler(tt.h)

		// 模拟请求
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8999/ss", nil)
		f(response, request)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expected (%d, %s), get (%d, %s)", tt.code, tt.message, response.Code, body)
		}
	}
}

func TestAppHandlerInServer(t *testing.T) {
	for _, tt := range tests {
		f := AppHandler(tt.h)

		// 真是请求
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.StatusCode != tt.code || body != tt.message {
			t.Errorf("expected (%d, %s), get (%d, %s)", tt.code, tt.message, response.StatusCode, body)
		}
	}
}
