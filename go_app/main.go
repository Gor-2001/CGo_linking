package main

/*
#cgo CFLAGS: -I../c_lib
#cgo LDFLAGS: -L../c_lib -lhelpers
#include "math_utils.h"
#include "string_utils.h"
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Add:", C.add(3, 4))
	fmt.Println("Multiply:", C.multiply(3, 4))
	fmt.Println("Length:", C.str_length(C.CString("Hello, world!")))
}
