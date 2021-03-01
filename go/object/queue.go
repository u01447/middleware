/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/3 下午3:09
 * Description:
 **/

package object

type Queue []int

func (q *Queue) Push(val int) error {
	*q = append(*q, val)
	return nil
}

func (q *Queue) Pop() (int, bool) {
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
