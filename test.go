package main
cgo CFLAGS: -fplugin=./test.so
import "C"
import "fmt"
typedef int (*intFunc) ();
int bridge_int_func(intFunc f)
 {
      return f();
 }
int fortytwo()
 {
      return 42;
 }
func main() {


    f := C.intFunc(C.fortytwo)
    fmt.Println(int(C.bridge_int_func(f)));
    // Output: 42
}
