/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/3 下午3:35
 * Description: 接口实现相关学习笔记-动物接口
 **/

package animal

import "fmt"

type AnimalsI interface {
	DuckI
	BehaviorI
	fmt.Stringer
}

func DuckBehavior(a AnimalsI) {
	name := a.GetName()
	dark := a.Shout("呱呱乱叫")
	fmt.Println(name, dark)
	fmt.Println(a.String())
}

func FindDuck(d DuckI) {
	name := d.GetName()
	fmt.Println(name)
}
