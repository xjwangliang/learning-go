/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 下午7:30
 * To change this template use File | Settings | File Templates.
 */
package mycmd

import (
	"os/exec"
	"os"
	"fmt"
)

func ExecCmd(){
	fmt.Printf("\n----------ExecCmd-----------\n")
	cmd := exec.Command("/bin/ls", "-l")  //Command(name string, arg ...string) *Cmd
	buf, err := cmd.Output()   //buf 是一个[]byte     (c *Cmd) Output() ([]byte, error)
	//err := cmd.Run()    //(c *Cmd) Run() error
	if err == nil {
		os.Stdout.Write(buf[:])
	}
	fmt.Printf("\n----------ExecCmd-----------\n")
}

