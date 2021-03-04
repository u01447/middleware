/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午1:05
 * Description:
 **/

package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type appHandler func(w http.ResponseWriter, request *http.Request) error

func AppHandler(app appHandler) func(w http.ResponseWriter, request *http.Request) {
	// 处理异常
	return func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			// panic 处理
			if err := recover(); err != nil {
				log.Printf("Unknown Error occured: %v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := app(w, request)

		if err != nil {
			log.Printf("Error occured: %s", err.Error())

			// 自定义异常处理
			if err, ok := err.(userError); ok {
				http.Error(w, err.Message(), http.StatusBadRequest)
				return
			}

			// 系统异常处理
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func UrlRequest(w http.ResponseWriter, request *http.Request) error {
	// 处理请求
	fmt.Println(request.URL.Path)
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("url path must be start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]

	if len(path) == 0 {
		return userError("url path illegal ")
	}

	var err error

	path, err = getCurrentPath(path)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}
