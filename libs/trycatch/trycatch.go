// author: kango2gler@gmail.com
// date: 2018-03-20

// try catch 实现包
package trycatch

//实现 try catch
func Trycatch(fun func(), handler func(interface{})) {
	defer func() {
		if err_ := recover(); err_ != nil {
			//err = errors.New(err_.(string))
			//fmt.Println("err:",reflect.TypeOf(err_))
			//if reflect.TypeOf(err).String() == "*errors.errorString" {
			//	// do something
			//}
			handler(err_)
		}
	}()
	fun()
}
