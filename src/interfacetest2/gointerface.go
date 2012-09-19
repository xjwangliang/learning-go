/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-19
 * Time: 上午11:03
 * To change this template use File | Settings | File Templates.
 */

	//http://www.ituring.com.cn/article/details/1339
package interfacetest2

import "fmt"

////自定义的File
//type File struct {
//	//...
//}
//
//func (f *File) Read(buf []byte) (n int, err error)
//func (f *File) Write(buf []byte) (n int, err error)
//func (f *File) Seek(off int64, whence int) (pos int64, err error)
//func (f *File) Close() error
//
//
//type IFile interface {
//	Read(buf []byte) (n int, err error)
//	Write(buf []byte) (n int, err error)
//	Seek(off int64, whence int) (pos int64, err error)
//	Close() error
//}
//
//type IReader interface {
//	Read(buf []byte) (n int, err error)
//}
//
//type IWriter interface {
//	Write(buf []byte) (n int, err error)
//}
//
//type ICloser interface {
//	Close() error
//}
//
//var file1 IFile = new(File)
//var file2 IReader = new(File)
//var file3 IWriter = new(File)
//var file4 ICloser = new(File)


type A interface {
	Area() float64
}
type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}
type Rect2 struct {
	x, y float64
	width, height float64
}

func (r Rect2) Area() float64 {
	return r.width * r.height
}

func Test(){
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	var rect5 Rect
	var rect6 *Rect
	//var rect57 = &Rect是错误的

	rect21 := new(Rect2)
	rect22 := &Rect2{}
	rect23 := &Rect2{0, 0, 100, 200}
	rect24 := &Rect2{width: 100, height: 200}
	var rect25 Rect2
	var rect26 *Rect2

	fmt.Printf("---------------------------------------------------")
	fmt.Printf("\nRect1 type:%T  value:%v",rect1,rect1)//Rect1 type:*interfacetest2.Rect  value:&{0 0 0 0}
	fmt.Printf("\nRect2 type:%T  value:%v",rect2,rect2)//Rect2 type:*interfacetest2.Rect  value:&{0 0 0 0}
	fmt.Printf("\nRect3 type:%T  value:%v",rect3,rect3)//Rect3 type:*interfacetest2.Rect  value:&{0 0 100 200}
	fmt.Printf("\nRect4 type:%T  value:%v",rect4,rect4)//Rect4 type:*interfacetest2.Rect  value:&{0 0 100 200}
	fmt.Printf("\nRect5 type:%T  value:%v",rect5,rect5)//Rect5 type:interfacetest2.Rect  value:{0 0 0 0}
	fmt.Printf("\nRect6 type:%T  value:%v",rect6,rect6)//Rect6 type:*interfacetest2.Rect  value:<nil>
	fmt.Printf("\n------------------------split---------------------------")
	fmt.Printf("\nRect21 type:%T  value:%v",rect21,rect21)//Rect21 type:*interfacetest2.Rect2  value:&{0 0 0 0}
	fmt.Printf("\nRect22 type:%T  value:%v",rect22,rect22)//Rect22 type:*interfacetest2.Rect2  value:&{0 0 0 0}
	fmt.Printf("\nRect23 type:%T  value:%v",rect23,rect23)//Rect23 type:*interfacetest2.Rect2  value:&{0 0 100 200}
	fmt.Printf("\nRect24 type:%T  value:%v",rect24,rect24)//Rect24 type:*interfacetest2.Rect2  value:&{0 0 100 200}
	fmt.Printf("\nRect25 type:%T  value:%v",rect25,rect25)//Rect25 type:interfacetest2.Rect2  value:{0 0 0 0}
	fmt.Printf("\nRect26 type:%T  value:%v",rect26,rect26)//Rect26 type:*interfacetest2.Rect2  value:<nil>

	fmt.Printf("\n---------------------------------------------------")

//	Rect1 type:*interfacetest2.Rect  value:&{0 0 0 0}
//	Rect2 type:*interfacetest2.Rect  value:&{0 0 0 0}
//	Rect3 type:*interfacetest2.Rect  value:&{0 0 100 200}
//	Rect4 type:*interfacetest2.Rect  value:&{0 0 100 200}
//	Rect5 type:interfacetest2.Rect  value:{0 0 0 0}
//	Rect6 type:*interfacetest2.Rect  value:<nil>

	//方法
	//通过接口：
	//接收者是指针不能传递值 正确 (没有相应的值接收者)
	//接受者是值能够传递指针 正确 (自动生成相应的指针接收者)

	//不通过接口：
	//接收者是指针能够传递值
	//接收者是值能够传递指针

	//指针var name *MyType使用之前必须初始化,不论MyType的接收者是指针还是值,否则是nil


	fmt.Printf("---------------------------------------------------")
	fmt.Printf("\n------------------------查询接口---------------------------")
   	var sup interface {} = rect1
	t1,ok1:=sup.(interface{}) //t:type *interfacetest2.Rect,value &{0 0 0 0}     	ok:type bool,value:true
	t2,ok2:=sup.(*interface{})//t:type *interface {},value <nil>     				ok:type bool,value:false
	t3,ok3:=sup.(Rect)   	  //t:type interfacetest2.Rect,value {0 0 0 0}     		ok:type bool,value:false
	t4,ok4:=sup.(*Rect)       //t:type *interfacetest2.Rect,value &{0 0 0 0}     	ok:type bool,value:true
	t5,ok5:=sup.(A)        	  //t:type *interfacetest2.Rect,value &{0 0 0 0}     	ok:type bool,value:true
	t6,ok6:=sup.(*A) 	      // t:type *interfacetest2.A,value <nil>     			ok:type bool,value:false
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t1,t1,ok1,ok1)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t2,t2,ok2,ok2)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t3,t3,ok3,ok3)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t4,t4,ok4,ok4)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t5,t5,ok5,ok5)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t6,t6,ok6,ok6)


	fmt.Printf("\n------------------------查询接口---------------------------")
	var sup2 interface {} = rect5
	t1,ok1 =sup2.(interface{}) //t:type interfacetest2.Rect,value {0 0 0 0}     		ok:type bool,value:true 	###
	t2,ok2 =sup2.(*interface{})//t:type *interface {},value <nil>     				ok:type bool,value:false
	t3,ok3 =sup2.(Rect)   	  //t:type interfacetest2.Rect,value {0 0 0 0}     		ok:type bool,value:true 	###
	t4,ok4 =sup2.(*Rect)       //t:type *interfacetest2.Rect,value <nil>     		ok:type bool,value:false 	###
	t5,ok5 =sup2.(A)        	  //t:type <nil>,value <nil>     						ok:type bool,value:false 	###(最大的不同)
	t6,ok6 =sup2.(*A) 	      // t:type *interfacetest2.A,value <nil>     			ok:type bool,value:false
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t1,t1,ok1,ok1)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t2,t2,ok2,ok2)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t3,t3,ok3,ok3)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t4,t4,ok4,ok4)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t5,t5,ok5,ok5)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t6,t6,ok6,ok6)
	fmt.Printf("\n---------------------------------------------------")


	fmt.Printf("---------------------------------------------------")
	fmt.Printf("\n------------------------查询接口---------------------------")
	sup  = rect21
	t21,ok21 :=sup.(interface{}) //t:type *interfacetest2.Rect2,value &{0 0 0 0}     ok:type bool,value:true
	t22,ok22 :=sup.(*interface{})//t:type *interface {},value <nil>     ok:type bool,value:false
	t23,ok23 :=sup.(Rect2)   	  //t:type interfacetest2.Rect2,value {0 0 0 0}     ok:type bool,value:false
	t24,ok24 :=sup.(*Rect2)       //t:type *interfacetest2.Rect2,value &{0 0 0 0}     ok:type bool,value:true
	t25,ok25 :=sup.(A)        	  //t:type *interfacetest2.Rect2,value &{0 0 0 0}     ok:type bool,value:true
	t26,ok26 :=sup.(*A) 	      //t:type *interfacetest2.A,value <nil>     ok:type bool,value:falseIs 5 even? false
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t21,t21,ok21,ok21)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t22,t22,ok22,ok22)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t23,t23,ok23,ok23)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t24,t24,ok24,ok24)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t25,t25,ok25,ok25)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t26,t26,ok26,ok26)


	fmt.Printf("---------------------------------------------------")
	fmt.Printf("\n------------------------查询接口---------------------------")
	sup  = rect25
	t21,ok21 =sup.(interface{}) //t:type interfacetest2.Rect2,value {0 0 0 0}     ok:type bool,value:true
	t22,ok22 =sup.(*interface{})//t:type *interface {},value <nil>     ok:type bool,value:false
	t23,ok23 =sup.(Rect2)   	  //t:type interfacetest2.Rect2,value {0 0 0 0}     ok:type bool,value:true
	t24,ok24 =sup.(*Rect2)       //t:type *interfacetest2.Rect2,value <nil>     ok:type bool,value:false
	t25,ok25 =sup.(A)        	  //t:type interfacetest2.Rect2,value {0 0 0 0}     ok:type bool,value:true
	t26,ok26 =sup.(*A) 	      //t:type *interfacetest2.A,value <nil>     ok:type bool,value:falseIs 5 even? false
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t21,t21,ok21,ok21)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t22,t22,ok22,ok22)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t23,t23,ok23,ok23)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t24,t24,ok24,ok24)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t25,t25,ok25,ok25)
	fmt.Printf("\nt,ok:=rect1.(interface{}) t:type %T,value %v     ok:type %T,value:%v",t26,t26,ok26,ok26)


	//小结：
	//只能用接口查询
	//大部分情况下，查询接口，成功返回实际类型和true，失败返回参数类型和false
	//查询是否是某个接口，加*号是不成功的，返回参数类型（值为nil）和false
	//不管指针还是值类型，查询interface{}接口（不是struct）都是成功的
	//实际类型是指针，查询struct需要加上*才会成功；同样实际类型是值，查询struct不能加*才会成功
	//查询接口，就比较复杂：
	//1　如果struct的实现方法的接收者是指针，而对象是struct指针，则成功
	//２ 如果struct的实现方法的接收者是值，不管对象是struct指针还是值，都成功
	//实际上是这个原理
	//------------->>>
	//方法
	//通过接口：
	//接收者是指针不能传递值 正确 (没有相应的值接收者)
	//接受者是值能够传递指针 正确 (自动生成相应的指针接收者)


	//所以如果一个struct实现了接口的方法，如果接收者是值，那么也就有了接收者是指针的方法，可以传递值和指针；
	//但是如果接收者是指针，没有接收者是值的方法，只能传递指针


	//(记住：查接口加*不成功;查interface{}永远成功;查struct类型匹配才成功;查接口还要看struct的接收者与对象的类型
	Test2()
	Test3()
	Test4()

}



func Test2(){
	fmt.Printf("\n------Test2-------")
	rect1 := new(Rect)
	fmt.Printf("\n1 rect1 value %v,area %v",rect1,rect1.Area())

	var sup interface {} = rect1
	sup.(*Rect).height = 10
	sup.(*Rect).width = 10

	fmt.Printf("\n1 rect1 value %v,area %v",rect1,sup.(*Rect).Area())


	ref := rect1
	ref.height = 20
	ref.width = 20
	fmt.Printf("\n1 ref value %v,area %v",rect1,rect1.Area())
	fmt.Printf("\n------Test2-------")

}


type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}
//测试接口赋值
func Test3(){
	var a Integer = 1
	var b LessAdder = &a     //... (1)
	//var c LessAdder = a      //... (2)   	Integer does not implement LessAdder (Add method requires pointer receiver)
	fmt.Printf("\n b %T %v",b,b)

}


type Base struct {
	//...
}

func (base *Base) Foo() {
	fmt.Printf(" Foo() ")
}
func (base *Base) Bar() {
	fmt.Printf("  Bar() ")
}

type Foo struct {
	Base
	age int
	//...
}

type Foo2 struct {
	*Base
	//...
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
	fmt.Printf("  Foo Bar() ")
}

func (base *Foo) Third() {
	fmt.Printf("  Third()  ")
}
//测试组合
func Test4(){
	var foo Foo

	fmt.Printf("\n>> foo type %T value %v",foo,foo)//foo type interfacetest2.Foo value {{} 0}

	fmt.Printf("\n>>")
	foo.Bar()        	// Bar()   Foo Bar()
	fmt.Printf("\n>>")
	foo.Base.Bar()      // Bar()
	fmt.Printf("\n>>")
	foo.Foo()          	//Foo()
	fmt.Printf("\n>>")
	foo.Base.Foo()     	//Foo()
	fmt.Printf("\n>>")


	var foo2 Foo2
	//foo2.Base = &Base{}
	fmt.Printf("\n>>")
	fmt.Printf("\n>> foo2 type %T value %v",foo2,foo2)    //foo2 type interfacetest2.Foo2 value {<nil>}
	foo2.Base.Foo()
	fmt.Printf("\n>>")

}
