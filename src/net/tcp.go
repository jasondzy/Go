package main 
import(
	"fmt"
	"os"
	"net"
	"bytes"
	"io"
)

func main(){
	if len(os.Args) != 2 {
		fmt.Println(" Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("tcp", service)

	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	os.Exit(1)
	}
}

func readFully(conn net.Conn) ([] byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for{ //此处为一直在这里进行循环的读取网络数据
		n, err := conn.Read(buf[0:]) //读取的数据存放到buf中，每次存放的大小事512字节
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF{
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}









