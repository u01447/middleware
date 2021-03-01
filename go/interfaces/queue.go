/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/3 下午5:55
 * Description:
 **/

package animal

type Queue []interface{}

func (q *Queue) Push(val interface{}) error {
	*q = append(*q, val)
	return nil
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, true
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
