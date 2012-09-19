/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 下午7:51
 * To change this template use File | Settings | File Templates.
 */
package myjson
import ("encoding/json"; "log"; "os")

//http://golangwiki.org/wiki/index.php?title=JSON%E5%92%8CGo&variant=zh
func TestJson() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		//if err := dec.Decode(&v); err != nil
		//json: Unmarshal(non-pointer map[string]interface {})
		if err := dec.Decode(&v); err != nil {//(dec *Decoder) Decode(v interface{}) error
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				v[k] = nil//, false
			}
		}
		//if err := enc.Encode(v); err != nil也可以，有什么区别？
		if err := enc.Encode(	v); err != nil {  //(enc *Encoder) Encode(v interface{}) error
			log.Println(err)
		}
	}
}
//由于Readers和Writers被广泛使用，因此Encoder和Decoder类型可被应用于许多场景，比如对HTTP连接、WebSocket或文件的读和写。

