package main

/*
#include <stdio.h>
#include <stdlib.h>
void hello_from_c(const char* name){
	printf("hello %s from C\n",name);
}
*/
import "C"
import "unsafe"

func main() {
	name := "Golang"
	cStr := C.CString(name)            // 类型转换成C字符数组
	defer C.free(unsafe.Pointer(cStr)) // 释放
	C.hello_from_c(cStr)               // 调用
}
