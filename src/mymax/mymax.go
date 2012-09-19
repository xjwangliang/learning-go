/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 下午4:54
 * To change this template use File | Settings | File Templates.
 */
package mymax

func Less(l, r interface{}) bool {
	switch l.(type) {
	case int:
		if _, ok := r.(int); ok {
			return l.(int) < r.(int)
		}
	case float32:
		if _, ok := r.(float32); ok {
			return l.(float32) < r.(float32)
		}
	}
	return false
}

func TestMax(){
	var a, b, c int = 5, 15, 0
	var x, y, z float32 = 5.4, 29.3, 0.0
	if c = a; Less(a, b) {
		c = b
	}
	if z = x; Less(x, y) {
		z = y
	}
	println("\n\n>>>>>>>>>>>>>max number is :\n")
	println(c, z)
}
