package main

// #include <stdio.h>
// #include <stdlib.h>
//
// // A static function in C is a function that has a scope that is limited to its object file
// // You can think of a static function as being "private" to its *.c file (although that is not strictly correct)
// static void myprint(char* s) {
//   printf("%s, in c language\n", s);
// }
//
// // intFunc is a typedef of a function pointer with no input, but returns int
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import (
	"fmt"
	"unsafe"

	point "github.com/dirkarnez/cgo-boilerplate/pointgo"
)

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))

	f := C.intFunc(C.fortytwo)             // 'grab' a c function **pointer** to `f`
	fmt.Println(int(C.bridge_int_func(f))) // put `f` to cgo's function signature
	// Output: 42

	p := point.NewPoint(0.0, 0.0)
	q := point.NewPoint(3.0, 4.0)

	dist := point.Distance(p, q)
	fmt.Println(dist)

	p.Delete()
	q.Delete()

	fmt.Scanln()
}
