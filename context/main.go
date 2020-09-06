package main

import (
	"net/http"

	"ogo"
)

func main() {
	r := ogo.New()
	r.GET("/", func(c *ogo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ogo</h1>")
	})
	r.GET("/hello", func(c *ogo.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ogo.Context) {
		c.JSON(http.StatusOK, ogo.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}