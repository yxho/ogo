package main

import (
	"fmt"
	"html/template"
	"net/http"
	"ogo"
	"time"
)

//func main() {
//	r := ogo.New()
//	r.GET("/", func(c *ogo.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	})
//
//	r.GET("/hello", func(c *ogo.Context) {
//		// expect /hello?name=geektutu
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//	})
//
//	r.GET("/hello/:name", func(c *ogo.Context) {
//		// expect /hello/geektutu
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//	})
//
//	r.GET("/assets/*filepath", func(c *ogo.Context) {
//		c.JSON(http.StatusOK, ogo.H{"filepath": c.Param("filepath")})
//	})
//
//	r.Run(":9999")
//}

//func main(){
//	r:=ogo.New()
//	r.GET("/index",func(c *ogo.Context){
//		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
//	})
//	v1:=r.Group("v1")
//	{
//		v1.GET("/", func(c *ogo.Context) {
//			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//		})
//
//		v1.GET("/hello", func(c *ogo.Context) {
//			// expect /hello?name=geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//		})
//	}
//	v2 := r.Group("/v2")
//	{
//		v2.GET("/hello/:name", func(c *ogo.Context) {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		})
//		v2.POST("/login", func(c *ogo.Context) {
//			c.JSON(http.StatusOK, ogo.H{
//				"username": c.PostForm("username"),
//				"password": c.PostForm("password"),
//			})
//		})
//
//	}
//
//	r.Run(":9999")
//}


//middlewares
//分组添加中间件 group.Use -> ServeHTTP根据前缀判断分组，中间件列表后，赋值给 c.handlers
//在路由的 handle中将从路由匹配得到的 Handler 添加到 c.handlers列表中，执行c.Next()。

//func onlyForV2() ogo.HandlerFunc {
//	return func(c *ogo.Context) {
//		// Start timer
//		t := time.Now()
//		// if a server error occurred
//		c.Fail(500, "Internal Server Error!")
//		// Calculate resolution time
//		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
//	}
//}
//
//func main() {
//	r := ogo.New()
//	r.Use(ogo.Logger()) // global midlleware
//	r.GET("/", func(c *ogo.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello ogo</h1>")
//	})
//
//	v2 := r.Group("/v2")
//	v2.Use(onlyForV2()) // v2 group middleware
//	{
//		v2.GET("/hello/:name", func(c *ogo.Context) {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		})
//	}
//
//	r.Run(":9999")
//}

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := ogo.New()
	r.Use(ogo.Logger())
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLGlob("D:\\WorkSpace\\GoProjects\\src\\ogo\\ogo_router\\templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *ogo.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *ogo.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", ogo.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *ogo.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", ogo.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}