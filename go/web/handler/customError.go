/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午4:43
 * Description:
 **/

package handler

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}
