package main
import (
	 "fmt"
	 "even"
	"interfacetest"
	"mysort"
	"mul2"
	"mymax"
	"myfile"
	"mycmd"
	"myjson"
	"interfacetest2"
	"mybinary"
	"data"
)

/* Print something */

//constant。它们在编译时被创建，只能是数字、字符串或布尔值； const x = 42 
const (
	a = iota	//a=0
	b = iota	//b=0
)
const (
	a1 = 0   //Is an int now
	b1 string = "0"
)


var (
	x int
	b2 bool
)

func evenTest() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
}


type NameAge struct {
	name string   //不导出
	age int   //不导出
}

func (in1 *NameAge) doSomething(in2 int) {in1.age = in2 }

func interfaceTest() {
	//var s = interfacetest.S(1)
	//s2 := interfacetest.S
	s := new(interfacetest.S)
	s2 :=  interfacetest.S(*s)  //   interfacetest.S类型，不能作为interfacetest.InterTest的参数

	s.Age = 99
	s.Put(88)

	s2.Put(199)
	s2.Age = 999
	fmt.Printf("interfaceTest\n")

	interfacetest.InterTest(s)
	//接收者是指针不能传递值 正确
	//接受者是值能够传递指针 ？
	interfacetest.InterTest(&s2)   //	interfacetest.S does not implement interfacetest.MyI (Get method requires pointer receiver)


	fmt.Printf("s type is  %T ,s2 type is %T\n",s,s2)     //s type is  *interfacetest.S ,s2 type is interfacetest.S

	var mys interfacetest.S
	mys.Age = 88
	interfacetest.InterTest(&mys) //	interfacetest.S does not implement interfacetest.MyI (Get method requires pointer receiver)
	mys.Get();
	mys.Put(9)

	var mysw *interfacetest.S

	//mysw.Age = 90     //panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Printf("mys type is  %T ,mysw type is %T value %v\n",mys,mysw ,mysw)//mys type is  interfacetest.S ,mysw type is *interfacetest.S
	//mysw.Put(89)  //会出错  panic: runtime error: invalid memory address or nil pointer dereference
	mysw = s
	mysw.Age = 90
	mysw.Get();
	interfacetest.InterTest(mysw)

	var mys2 interfacetest.S2
	mys2.Age = 98
	interfacetest.InterTest(mys2)
	mys2.Get()

	fmt.Printf("mys type is  %T ,mys2 type is %T value %v\n",mys,mys2 ,mys2)  //mys type is  interfacetest.S ,mys2 type is interfacetest.S2

	var mys2w *interfacetest.S2
	//mys2w.Age = 90 panic: runtime error: invalid memory address or nil pointer dereference
	//mys2w.Get() 会出错
	mys2w = &mys2
	mys2w.Age = 90
	interfacetest.InterTest(mys2w)
	mys2w.Get()
	fmt.Printf("mys2 type is  %T ,mys2w type is %T\n",mys2,mys2w )   //mys2 type is  interfacetest.S2 ,mys2w type is *interfacetest.S2

	//t,ok := mys.(interfacetest.MyI)   //invalid type assertion: mys.(interfacetest.MyI) (non-interface type interfacetest.S on left)
	var myI interfacetest.MyI
	//myI.Put(199)    出错
	//var myI2 interfacetest.S2
	t,ok := myI.(interface{})

//	if t, ok := mys.(interfacetest.MyI); ok {
//		// 对于某些实现了接口I 的
//		// t 是其所拥有的类型
//	}

	fmt.Printf("myI type is  %T ,myI type is ok %v\n",t,ok )//myI type is  <nil> ,myI type is ok false

	t,ok = myI.(interfacetest.S2)
	fmt.Printf("myI type is  %T ,myI type is ok %v\n",t,ok ) //myI type is  interfacetest.S2 ,myI type is ok false
	//a,b := s.(interfacetest.MyI)
	t,ok = myI.(interfacetest.MyI)
	fmt.Printf("myI type is  %T ,myI type is ok %v\n",t,ok )  //myI type is  <nil> ,myI type is ok false


	var ddd interface{}
	ddd = s
	t,ok = ddd.(interfacetest.MyI)
	fmt.Printf("%T ddd type is  %T ,myI type is ok %v\n",ddd,t,ok )  //*interfacetest.S ddd type is  *interfacetest.S ,myI type is ok true

	ddd = mys
	t,ok = ddd.(interfacetest.MyI)
	fmt.Printf("%T ddd type is  %T ,myI type is ok %v\n",ddd,t,ok )  //interfacetest.S ddd type is  <nil> ,myI type is ok false


	var n *NameAge     //在NameAge中的doSomething中用到了成员就会报错  anic: runtime error: invalid memory address or nil pointer dereference
	n = &NameAge{"name",2}
	n.doSomething(2)


	var n2 NameAge
	n2.doSomething(2)
	fmt.Printf("\nn2 type is  %T , value is  %v\n",n2,n2 )
	fmt.Printf("\nn type is  %T , value is  %v\n",n,n )

	var rtest *interfacetest.AA
	rtest.Get()
	fmt.Printf("\nrtest type is  %T , value is  %v\n",rtest,rtest )//rtest type is  *interfacetest.AA , value is  <nil>

	//下面有错误，上面没有错误
//	var ftest *interfacetest.S
//	ftest.Get()

}

func mulTest(){
	const (
		Enone = 2
		Einval = 5
	)
	ar := [...]string {Enone: "no error", Einval: "invalid argument"}    //对于数组，5是最大的数字，所以是len是6
	sl := []string {Enone: "no error", Einval: "invalid argument"}
	ma := map[int]string{Enone: "no error", Einval: "invalid argument"}
	fmt.Printf("ar type is  %T ,value is %v\n",ar,ar )   //ar type is  [6]string ,value is [  no error   invalid argument]
	fmt.Printf("sl type is  %T ,value is %v\n",sl,sl )   //sl type is  []string ,value is [  no error   invalid argument]
	fmt.Printf("ma type is  %T ,value is %v\n",ma,ma )   //ma type is  map[int]string ,value is map[5:invalid argument 2:no error]
}


func mulTest2(){

	fmt.Printf("\n--mulTest2----mulTest2--\n")


	m := []mul2.E{1, 2, 3, 4}
	s := []mul2.E{"a", "b", "c", "d"}
	mf := mul2.Map(m, mul2.Mult2)
	sf := mul2.Map(s, mul2.Mult2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

type Foo int
func (self Foo) Emit() {
	fmt.Printf("\n--Emit-- %v --Emit--\n", self)
}
type Emitter interface {
	Emit()
}

func methodTest2(e Emitter){
	e.Emit()
}
func methodTest(){
	var f Foo
	f.Emit()

	methodTest2(f)
}

//sort测试接口
func testSort(param []int){   //use of [...] array outside of array literal

	fmt.Printf("paramater type is  %T ,p item type %v,value is %v\n",param,param )

	ints := mysort.Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := mysort.Xs{"nut", "ape", "elephant", "zoo", "go"}
	mysort.Sort(ints)
	fmt.Printf("%v\n", ints)
	mysort.Sort(strings)  //系统也有sort
	fmt.Printf("%v\n", strings)
}


type Interface interface {
	mysort.Sorter
	Push(x interface{})
	Pop() interface{}
}




func g(some interface{}) int {
	fmt.Printf("g type is %T   ",some)
	v, ok := some.(interfacetest.MyI)
	fmt.Printf("g2 type is %T  %v is OK ? %v ",v,v,ok)
	if ok { // 检查是否可以转换
		return v.Get() // 如果可以，调用Get()
	}
	return -1 // 随便返回个什么
}

func interfaceTest2(){
	var s =  new (interfacetest.S)
	var s2 interfacetest.S
	s2.Put(22)
	s.Put(100)
	fmt.Println("\n------------interface test---------\n")


	fmt.Println(g(s))
	fmt.Println(g(s2))
	fmt.Println(g(&s2))
	i := 5
	fmt.Println(g(i))
	fmt.Printf(" type is %T  %T   %T ",s,s2,i)
//	g type is *interfacetest.S   g2 type is *interfacetest.S  &{100}  100
//	g type is interfacetest.S   g2 type is <nil>  <nil>  -1
//	g type is *interfacetest.S   g2 type is *interfacetest.S  &{22}  22
//	g type is int   g2 type is <nil>  <nil>  -1
//	type is *interfacetest.S  interfacetest.S   int
	fmt.Printf("\n------------interface test---------\n")
}

/**
var i int
	i = 15 //i := 15


	var b1 bool
	b1 = false//b:=false

	var j int32 //int8，int16，int32，int64 和byte，uint8，uint16，uint32，uint64。
	//byte 是uint8 的别名。浮点类型的值有float32 和float64 （没有float 类型）。64 位的整数和浮点数总是64 位的，即便是在32 位的架构上
 **/




// non-declaration statement outside function body在函数外只能声明，不能定义赋值
//  b redeclared in this block
//  x declared and not used在函数内声明变量但是没有使用


func TestMax(){
	var a, b, c int = 5, 15, 0
	var x, y, z float32 = 5.4, 29.3, 0.0
	if c = a; mymax.Less(a, b) {
		c = b
	}
	if z = x; mymax.Less(x, y) {
		z = y
	}
	fmt.Println("\n\n>>>>>>>>>>>>>max number is :\n")
	fmt.Println(c, z) //println(c,z)
	//fmt.Println(10,9.0)

}
func main() { 
	fmt.Printf("Hello, world\n")

	data.Test()
	data.Test2()
	mybinary.Test()
	interfacetest2.Test()
	evenTest()
	interfaceTest()
	mulTest()
	methodTest()

	testSort([]int {2,8,9})
	interfaceTest2()
	mulTest2()


	TestMax()

	myfile.TestFile()
	myfile.TestBufferedFile()
	myfile.TestReadLine()
	//myfile.TestHttpGet()


	mycmd.ExecCmd()
	myjson.TestJson()




	var i int
	i = 15 //i := 15

	fmt.Println(i)
	i2 := 10		//no new variables on left side of :=(不能重复)
	fmt.Println(i2)
	//
	//
	//
	//字符串是不可变的
	//var s string = "hello"
	//s[0] = 'c'   修改第一个字符为'c'，这会报错
	s := "hello"
	c := []rune(s)

	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)//cello


	//数组倒序
	var tt = []int{1,29,3,41,5}
    for i,j:=0,len(tt)-1; i < j; i,j=i+1,j-1 {
        tt[i],tt[j] = tt[j],tt[i]
    }
    fmt.Println(tt);


    tt1:= make([]int ,5)//被填充的{0,0,0,0,0}
    tt2:=make([]int ,0,5)//这时候tt的最大长度是5，但结果是{}

     fmt.Println(tt1);
     fmt.Println(tt2);

}
