package main

/*
#include <stdio.h>
#include <stdlib.h>

int myPrints(int i)
{
  return i + 10;
}
*/
import "C"
import (
	"fmt"
	//"unsafe"
)

func main() {
	cInt := C.int(1)
	a := C.myPrints(cInt)
	//if a == (*C.char)(unsafe.Pointer(uintptr(0))) {
	//	return
	//}

	//defer C.free(unsafe.Pointer(a))
	//fmt.Println("GoPrintln:", C.GoString(a))
	fmt.Println("Go output:", int(a))

}
