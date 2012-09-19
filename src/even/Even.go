/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 上午10:36
 * To change this template use File | Settings | File Templates.
 */
package even   //开始自定义的包
func Even(i int) bool {   //可导出函数
return i % 2 == 0
}
func odd(i int) bool {   //私有函数
	return i % 2 == 1
}

