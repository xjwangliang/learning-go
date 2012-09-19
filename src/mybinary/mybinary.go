/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-19
 * Time: 下午4:58
 * To change this template use File | Settings | File Templates.
 */
package mybinary

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

func Test() {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(61374),
		int8(-54),
		uint8(254),
	}
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	fmt.Printf("\n mybinary test %x", buf.Bytes())
}
