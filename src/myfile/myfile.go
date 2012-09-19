package myfile

import (
	"os"
	"fmt"
	"bufio"
	"net/http"
	"io/ioutil"
)

func TestFile() {
	fmt.Printf("\n----------Test File-----------\n")
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/wangliang/Desktop/笔记/host.txt")
	defer f.Close() //
	for {
		n, _ := f.Read(buf)
		if n == 0 { break }
		os.Stdout.Write(buf[:n])
	}
	fmt.Printf("\n----------Test File-----------\n")
}

func TestBufferedFile() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/wangliang/Desktop/笔记/host.txt")
	defer f.Close()
	r := bufio.NewReader(f)   //NewReader(rd io.Reader) *Reader,File有Read方法，所以实现了Reader接口
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 { break }
		w.Write(buf[0:n])
	}
}
func TestReadLine() {
	fmt.Printf("\n----------TestReadLine-----------\n")
	f, _ := os.Open("/Users/wangliang/Desktop/笔记/host.txt")
	defer f.Close()
	r := bufio.NewReader(f)   //NewReader(rd io.Reader) *Reader,File有Read方法，所以实现了Reader接口
	var index int = 0
	for  {
		s, ok := r.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Printf("line %d :%s",index,s)
		index++
	}
	fmt.Printf("\n----------TestReadLine-----------\n")

//		if t, ok := 2,true; ok {
//			// 对于某些实现了接口I 的
//			// t 是其所拥有的类型
//		}

	//s, ok := r.ReadString('\n')   //从输入中读取一行,s 保存了字符串，通过string 包就可以解析它|




}

func TestHttpGet() {
	fmt.Printf("\n----------TestHttpGet-----------\n")
	r, err := http.Get("http://www.google.com/?q=wangliang") //Get(url string) (r *Response, err error)
	if err != nil {
		fmt.Printf("%s\n", err.Error());
		return
	}
	//ReadAll(r io.Reader) ([]byte, error)
	b, err := ioutil.ReadAll(r.Body)  //Body io.ReadCloser,而它包含Reader和Closer，所以实现了Reader
	r.Body.Close()
	if err == nil { fmt.Printf("%s", string(b)) } //4.
	fmt.Printf("\n----------TestHttpGet-----------\n")
}
