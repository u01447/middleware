/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/2/19 下午5:16
 * Description: 接口实现相关学习笔记-鸭子接口实现
 **/

package mock

import "fmt"

type Duck struct {
	name string
	bark string
}

func (d *Duck) GetName() string {
	if d.name != "" {
		return d.name
	} else {
		return "这是一个鸭子！"
	}
}

func (d *Duck) Shout(dark string) string {
	if d.bark == "" {
		return "呱呱呱呱的叫"
	} else {
		return dark
	}
}

func (d *Duck) String() string {
	return fmt.Sprintf("Duck: { name = %s, bark = %s }", d.name, d.bark)
}
