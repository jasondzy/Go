package main 

import (
    "fmt"
    "log"
    "net/http"
    "sort"
    "strconv"
    "strings"
    "math"
)

const (
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
    form       = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
)

type statistics struct {
    numbers []float64
    mean    float64
    median  float64
    mode    []float64
    div     float64
}

func main(){

	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe("0.0.0.0:9001",nil);err != nil {
		log.Fatal("failed to start server",err)
	}
}

func homePage(writer http.ResponseWriter, request *http.Request){ //这个函数是固定的定义方式

	err := request.ParseForm()
	fmt.Fprintln(writer, pageTop, form) //这里的作用是将pageTop和form写入writer接口中去
	if err != nil {
		fmt.Fprint(writer, anError,err)
	}else {
		if numbers,message,ok := processRequest(request); ok{
			stats := getStats(numbers)
			fmt.Fprint(writer,formatStats(stats))
		} else if message != ""{
			fmt.Fprint(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
    var numbers []float64
    if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
        text := strings.Replace(slice[0], ",", " ", -1) //这里的作用是将切片中的，号替换成空格
        for _, field := range strings.Fields(text) { //遍历这个字符串，Fields函数默认采用的是空格的方式进行区分，所以field中存放的是按照空格分开的数据
            if x, err := strconv.ParseFloat(field, 64); err != nil { //这里将字符串形式的数字转换成float64类型
                return numbers, "'" + field + "' is invalid", false
            } else {
                numbers = append(numbers, x) //将转换后的值x放入numbers数组中去
            }
        }
    }
    if len(numbers) == 0 {
        return numbers, "", false // no data first time form is shown
    }
    return numbers, "", true //返回3个值
}

func formatStats(stats statistics) string {
    return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
<tr><td>Mode</td><td>%v</td></tr>
<tr><td>Div</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.mode, stats.div)
}

func getStats(numbers []float64) (stats statistics) {
    stats.numbers = numbers
    sort.Float64s(stats.numbers) //这里进行排序操作，
    stats.mean = sum(numbers) / float64(len(numbers)) //获取切片的中间值
    stats.median = median(numbers)
    stats.mode = mode(numbers)
    stats.div = div(numbers)
    return stats
}

func sum(numbers []float64) (total float64) {
    for _, x := range numbers { //这里使用的是用range的方式获取切片中的每一个值，然后再进行求和处理
        total += x
    }
    return total
}

func median(numbers []float64) float64 {
    middle := len(numbers) / 2
    result := numbers[middle]
    if len(numbers)%2 == 0 { //这里是对于偶数切片求平均值的方式
        result = (result + numbers[middle-1]) / 2
    }
    return result
}

func mode(numbers []float64) []float64 {
    var a []float64 = make([]float64,1,len(numbers))
    var b int = 0 
    for i,x := range numbers{
        count := 1
        temp := numbers[i:]
        fmt.Println(temp)
        for _,y := range temp{
            if x==y{
                count += 1
                fmt.Printf("count:%d\n",count)
            }
        }
        if b < count{
            a = a[:1] //这就是go中slice的删除操作，这个操作的作用是只保留了a的第一个元素
            a[0] = x
            b = count
        } else if b == count{
            a = append(a,x)
        }
    }
    fmt.Println(a)
    if len(a) == len(numbers){
        fmt.Printf("something worng happen")
        a := make([]float64,1,len(numbers))
        return a
    } else {
        return a
    }
}

func div(numbers []float64) (result float64) {
    sum := float64(0)
    avg := float64(0)
    for _, x := range numbers {
        sum += x
    }
    avg = sum/float64(len(numbers))
    sum = 0
    for _, y := range numbers{
        sum += (y-avg)*(y-avg)
    }

    result = math.Sqrt(sum/float64(len(numbers)-1))
    return result
}