package util

import "fmt"

// 自定义错误处理
func Error(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
	}
}
