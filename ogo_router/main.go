package main

import (
	"net/http"
	"ogo"
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

func main(){
	r:=ogo.New()
	r.GET("/index",func(c *ogo.Context){
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1:=r.Group("v1")
	{
		v1.GET("/", func(c *ogo.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *ogo.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ogo.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ogo.Context) {
			c.JSON(http.StatusOK, ogo.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}