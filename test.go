package main

import "C"

//export HelloWorld
func HelloWorld() *C.char {
    return C.CString("Hello from Go!")
}

func main() {} 
