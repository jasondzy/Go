package main 

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
)


var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readlines(infile string) (values []int, err error){
	file, err := os.Open(infile)
	if err != nil{
		fmt.Println("open file wrong")
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for{
		// line, isPrefix, err1 := br.ReadLine()
		line, err1 := br.ReadBytes('\n') //库函数中建议使用的是readbytes来进行处理，具体作用可见go的库函数解析
		if err1 != nil{
			if err1 != io.EOF{
				err = err1
			}
			break
		}
	
		// if isPrefix {
		// 	fmt.Println("A too long line , seems unexpected")
		// 	return
		// }

		str := string(line[:len(line)-1]) //由于readbytes读取的时候会读取每一行结尾的换行符，所以这里要排除结尾的换行符
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}

	return


}

func writelines(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Create file error")
		return err
	}

	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
	
}


func main() {
	flag.Parse()

	if infile != nil{
		fmt.Println("infile:", *infile, "outfile:", *outfile, "algorithm", *algorithm)
	}

	values, err := readlines(*infile)

	if err != nil {
		fmt.Println(" readlines failed")
		fmt.Println(err)
		return
	}

	fmt.Println("values:", values)

	writelines(values, *outfile)

}