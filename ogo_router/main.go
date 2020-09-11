package main

import (
	"net/http"
	"ogo"
)

func main() {
	r := ogo.New()
	r.GET("/", func(c *ogo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *ogo.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *ogo.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *ogo.Context) {
		c.JSON(http.StatusOK, ogo.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}