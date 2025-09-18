cgo-boilerplate
===============
This boilerplate demostrate golang compilation along with a C++ static library which has a C language public header interface

### Notes
- settings in `go build` command as in `local-build.cmd`

### TODOs
- [ ] github action (cross compilation)

### Tutorials
- [CGO基础 - Go语言高级编程](https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-02-basic.html)
  - ```go
    package main
    /*
    #include <stdio.h>
    
    void printint(int v) {
        printf("printint: %d\n", v);
    }
    */
    import "C"
    
    func main() {
    	v := 42
    	C.printint(C.int(v))
        
    }
    ```
### Reference
- [krumberg/cmake_go: CMake support for building Go programs that links to C/C++](https://github.com/krumberg/cmake_go)
