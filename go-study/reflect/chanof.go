package main

import (
	"fmt"
	"reflect"
)

// 函数说明  func ChanOf(dir ChanDir, t Type) Type
// 参数 : dir ChanDir 信道的方向. 类型是reflect.ChanDir  t Type 信道的类型 类型是reflect.Type
// 返回值: Type 返回的是 reflect.Type 类型
// 功能说明:  Chanof 返回信道类型与给定的方向和元素类型. 例如, 如果t是整型, 纳闷ChanOf(RecvDir, t) 这样就表示 <-chan int
// GC 运行时强加限制 64KB 的通道元素类型. 如果t的大小等于或超过此限制, ChanOf恐慌(Panic)

func main() {
	var i int
	var recvDir reflect.ChanDir = reflect.RecvDir
	var chanOf reflect.Type = reflect.ChanOf(recvDir, reflect.TypeOf(i))
	fmt.Println(chanOf.Kind())    // chan
	fmt.Println(chanOf.ChanDir()) // <-chan
	fmt.Println(chanOf.String())  // <-chan int

	var i1 int
	var recvDir1 reflect.ChanDir = reflect.SendDir
	var chanOf1 reflect.Type = reflect.ChanOf(recvDir1, reflect.TypeOf(i1))
	fmt.Println(chanOf1.Kind(), chanOf1.ChanDir(), chanOf1.String())
	// chan chan<- chan<- int

	var i2 int
	var recvDir2 reflect.ChanDir = reflect.BothDir
	var chanOf2 reflect.Type = reflect.ChanOf(recvDir2, reflect.TypeOf(i2))
	fmt.Println(chanOf2.Kind(), chanOf2.ChanDir(), chanOf2.String())
	// chan chan chan int

	var i3 string
	var recvDir3 reflect.ChanDir = reflect.BothDir
	var chanOf3 reflect.Type = reflect.ChanOf(recvDir3, reflect.TypeOf(i3))
	fmt.Println(chanOf3.Kind(), chanOf3.ChanDir(), chanOf3.String())
	// chan chan chan string
}
