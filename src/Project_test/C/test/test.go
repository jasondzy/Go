package main

//void SayHello(const char *s);
//extern char* ss;
//extern char buf[10];
/*struct info {
    char* name;
    int age;
};*/
import "C"
import "fmt"
import "reflect"
import "unsafe"

func main() {
	
	//这里的示例是如何通过reflect来获取C中的字符串	
	var s0 string
    s0Hdr := (*reflect.StringHeader)(unsafe.Pointer(&s0))
    s0Hdr.Data = uintptr(unsafe.Pointer(C.ss))
    s0Hdr.Len = 10

	fmt.Println("begin:",s0) //切记：这里输出的是s0而不是soHDR,因为s0Hdr只是一个存在指向s0的reflect.StringHeader的类型，并不是string类型，所以直接输出是有问题的
	

	//这里进行的是C语言中的struct的引用，Go中只能引用C语言中的struct定义，而不能引用结构体中的值，同时改变变量需要在上边进行导出
	//这里就是通过C语言中的struct进行变量的定义
	var ff C.struct_info
	//其中name仍然是C语言中的类型，所以这里要将GO语言中的字符串进行转换
	ff.name = C.CString("hello struct")
	//由于name是Cgo中的类型，所以再GO中进行输出的时候要进行转换
	fmt.Println(C.GoString(ff.name))
	// struct end

	//如下直接调用C语言中的变量，C.ss变量的类型是cgo中的类型
	fmt.Println(C.ss)
	fmt.Printf("%T\n", C.ss)

	//如下直接引用的是C语言中的数组变量，这里可有for var := range C.buf的方式对C.buf进行迭代操作
	fmt.Println(C.buf)
	fmt.Printf("%T", C.buf)

	//这里调用的是C语言中的函数
	C.SayHello(C.CString("hello test for cgo"))
}

//export pp
func pp(s *C.char) {
	fmt.Println(C.GoString(s))
}

