//注意:在Golang中，对于单元测试程序来说通常会有一些重要约束，主要如下:
//
//单元测试文件名必须为xxx_test.go(其中xxx为业务逻辑程序)
//
//单元测试的函数名必须为Testxxx(xxx可用来识别业务逻辑函数)
//
//单元测试函数参数必须为t *testing.T(测试框架强要求)
//
//测试程序和被测试程序文件在一个包package中
//
//运行在该目录 go test
//参数 t 上的 Fatal 和 Fatalf 方法被用于记录致命的被程序实体的状态错误。所谓致命错误是指使得测试无法继续进行的错误。例如：
//
//if listener == nil {
//t.Fatalf("Listener startup failing! (addr=%s)!\n", serverAddr)
//}

//调用 t.Fatal 方法相当于先后对 t.Log 和 t.FailNow 方法进行调用，而调用 t.Fatalf 方法则相当于先后对 t.Logf 和 t.FailNow 方法进行调用。
package ogo

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := (n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt")
	if !ok1 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be file1.txt")
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}

}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes shoule be 4")
	}
}
