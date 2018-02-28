package main 
import(
	"fmt"
	"net/http"
)

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/test/", test)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w,"hello world")
}

func test(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Location","https://www.baidu.com")
	w.WriteHeader(302) 
	// h := request.Header
	// fmt.Fprintln(w, h)
}
