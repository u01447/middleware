/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午12:32
 * Description:
 **/

package web

import (
	"net/http"
	"offer/note/web/handler"
)

func Server() {
	http.HandleFunc("/", handler.AppHandler(handler.UrlRequest))
	err := http.ListenAndServe("127.0.0.1:8999", nil)
	if err != nil {
		panic(err)
	}
}
