/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 上午10:30
 * To change this template use File | Settings | File Templates.
 */
package interfacetest

import "fmt"

type S struct { Age int }
func (p *S) Get() int { return p.Age }
func (p *S) Put(v int) { p.Age = v }

type S2 struct { Age int }
func (p S2) Get() int { return p.Age }
func (p S2) Put(v int) { p.Age = v}

type MyI interface {
	Get() int
	Put(int)
}


type AA struct { Age int }
func (p *AA) Get() int { return 10 }

func InterTest(p MyI) {
	fmt.Println(p.Get())
	p.Put(1)
}

//获取s 的地址，而不是S 的值的原因，是因为在s 的指针上定义了方法，参阅
//上面的代码6.1。这并不是必须的——可以定义让方法接受值——但是这样的话
//Put 方法就不会像期望的那样工作
