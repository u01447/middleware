/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午1:46
 * Description:
 **/

package handler

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

func getCurrentPath(base string) (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return getParentDirectory(filepath.Dir(file)) + "source/" + base, nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/")+1)
}
