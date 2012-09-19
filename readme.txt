func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0

    return f
}
有许多冗长的内容。可以使用复合声明使其更加简洁，每次只用一个表达式创建一个新的实例。
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}   //Create a new File
    return &f                     //Return the address of f
}
返回本地变量的地址没有问题；在函数返回后，相关的存储区域仍然存在。

事实上，从复合声明获取分配的实例的地址更好，因此可以最终将两行缩短到一行。(从复合声明中获取地址，意味着告诉编译器在堆中分配空间，而不是栈中。)

    return &File{fd, name, nil, 0}
The items (called of a composite +literal are laid out in order and must all be 所有的项目（称作字段）都必须按顺序全部写上。然而，通过对元素用字段: 值成对的标
识，初始化内容可以按任意顺序出现，并且可以省略初始化为零值的字段。因此可以这样

    return &File{fd: fd, name: name}

在特定的情况下，如果复合声明不包含任何字段，它创建特定类型的零值。表达式new(File) 和&File{} 是等价的。


复合声明同样可以用于创建array，slice 和map，通过指定适当的索引和map 键来标识字段。在这个例子中，无论是Enone，Eio 还是Einval 初始化都能很好的
工作，只要确保它们不同就好了。

ar := [...]string {Enone: "no error", Einval: "invalid argument"}
sl := []string {Enone: "no error", Einval: "invalid argument"}
ma := map[int]string{Enone: "no error", Einval: "invalid argument"}




//t,ok := mys.(interfacetest.MyI)   //invalid type assertion: mys.(interfacetest.MyI) (non-interface type interfacetest.S on left)
if t, ok := mys.(interfacetest.MyI); ok {
		// 对于某些实现了接口I 的
		// t 是其所拥有的类型
	}

是有错误的



var mys interfacetest.S
	mys.Age = 88
	interfacetest.InterTest(&mys)
	mys.Get();
	var mysw *interfacetest.S
	//mysw.Age = 90
	fmt.Printf("mys type is  %T ,mysw type is %T\n",mys,mysw )//mys type is  interfacetest.S ,mysw type is *interfacetest.S
	//mysw.Put(89)  会出错