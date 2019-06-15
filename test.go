package main
import "C"
import "fmt"

func main() {


    var f := C.intFunc(C.fortytwo)
    fmt.Println(int(C.bridge_int_func(f)));
    // Output: 42
}
//cgo CFLAGS: -fplugin=./test.so

int 
bridge_int_func(intFunc f)
 {
      return f();
 }
int fortytwo()
 {
      return 42;
 }
typedef int (*intFunc) ();
