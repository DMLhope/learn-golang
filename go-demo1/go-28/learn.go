//错误处理
//go语言允许函数和方法同时返回多个值
//按照管理，函数在返回错误时，最后面的返回值应该用来表示错误
//调用函数后应立即检测是否发生错误
//	如果没有错误发生，那么返回的错误值为nil
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./sdjfla")
	if err != nil {
		fmt.Println(err)
		//传入非0值表示错误退出
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "...")
	if err != nil {
		f.Close()
		return err
	}
	f.Close()
	return err
}
