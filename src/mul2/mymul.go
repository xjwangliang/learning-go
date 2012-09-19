/**
 * Created with IntelliJ IDEA.
 * User: wangliang
 * Date: 12-9-18
 * Time: 下午2:26
 * To change this template use File | Settings | File Templates.
 */
package mul2

//* define the empty interface as a type
type E interface{}
func Mult2(f E) E {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(
		string)
	}
	return f
}
func Map(n []E, f func(E) E) []E {
	m := make([]E, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

