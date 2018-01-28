package main 

import (
	// "fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, req *http.Request){ //这里是服务器在请求来临的时候的处理函数，req是接收的参数，w是相应的参数，这个函数的目的就是填充相应参数
	// fmt.Println(" connect sucess ")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("This is an example server.\n"))
}

func main(){

	http.HandleFunc("/test", testHandle)

	http.ListenAndServe("0.0.0.0:8080", nil)
}