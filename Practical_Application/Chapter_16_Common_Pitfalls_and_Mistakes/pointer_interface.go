package main

import (
	"fmt"
)

type nexter interface {
	next() byte
}

func nextFew1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i < num; i++ {
		b[i] = n.next()
	}
	return b
}
// func nextFew2(n *nexter, num int) []byte {
// 	var b []byte
// 	for i := 0; i < num; i++ {
// 		b[i] = n.next() // 编译错误：n.next 未定义（*nexter 类型没有 next 成员或 next 方法）
// 	}
// 	return b
// }
func main() {
	fmt.Println("Hello World!")
}
