/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/4 下午4:39
 * Description:
 **/

package errs

type UserError interface {
	error
	Message() string
}
