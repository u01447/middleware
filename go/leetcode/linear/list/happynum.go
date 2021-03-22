/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/3/5 上午10:56
 * Description:
 **/

package list

func IsHappy(n int) bool {
	p, q := getNext(n), getNext(getNext(n))
	for p != q && q != 1 {
		p = getNext(p)
		q = getNext(getNext(q))
	}
	return q == 1
}

func getNext(n int) int {
	x := 0
	for n > 0 {
		x += (n % 10) * (n % 10)
		n = n / 10
	}
	return x
}
