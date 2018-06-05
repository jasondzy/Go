package service

import (
	. "github.com/smartystreets/goconvey/convey" //注意：这里的.是一个import技巧，表示的是convey的别名，这样在下边的代码中不用使用convey.xxx了
	"testing"
	"net/http/httptest"
	"fmt"
)

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		
		req := httptest.NewRequest("GET", "/invalid/123", nil)  //这里使用httptest包中的函数创建一个request请求
		resp := httptest.NewRecorder()   //创建一个recorder类型，该类型实现了ResponseWriter接口，所以可以作为下边的ServeHTTP函数的参数使用

		Convey("Then the response should be a 404", func() {
			NewRouter().ServeHTTP(resp, req)  //这里调用mux中router的serverHTTP方法来填充resp这个结构体，这里直接调用serverHTTP这个函数来模仿一个HTTP的请求，该请求就是req
			fmt.Println("the value of resp.Code:",resp.Code)
			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404) //这里直接使用So函数的原因就是在import中使用了.这个别名进行导入的原因
			})
		})

	})
}