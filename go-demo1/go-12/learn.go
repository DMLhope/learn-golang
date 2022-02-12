package go12

func main() {

}

// 常规函数定义
// 在Go里，大写字母开头的函数、变量或其他标识符都会被导出对其他包可用
// 小写开头的不行
func Test(num1, num2 int64) {

}

// 常规int64返回函数定义
func test(num1, num2 int64) int64 {

}

// 多返回值函数定义
func test1(num1, num2 int64) (i int64, err error) {

}

// 函数返回值名字可以省略
func test2(num1, num2 int64) (int64, error) {

}

// 可变参数函数
func Println(a ...interface{}) (n int, err error) {
	// ...表示函数的参数数量可变a
	// 参数a的类型位interface{},是一个空接口
}
