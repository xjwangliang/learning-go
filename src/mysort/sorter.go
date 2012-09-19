/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 下午12:23
 * To change this template use File | Settings | File Templates.
 */
package mysort

type Sorter interface {
	Len() int   //len() 作为方法
	Less(i, j int) bool   //p[j] < p[i] 作为方法
	Swap(i, j int)   //p[i], p[j] = p[j], p[i] 作为方法
}

type Xi []int
type Xs []string

//处理整数的
func (p Xi) Len() int { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int) { p[i], p[j] = p[j], p[i] }


//处理字符串的：
func (p Xs) Len() int { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int) { p[i], p[j] = p[j], p[i] }


func Sort(x Sorter) {
	for i := 0; i < x.Len() - 1; i++ { //1.
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}
