package main

import (
	"fmt"
	"gee/gee"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func middleWaresTestForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	// r := gee.New()
	// r.Use(gee.Logger())
	// r.SetFuncMap(template.FuncMap{
	// 	"FormatAsDate": FormatAsDate,
	// })
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/assets", "./static")

	// stu1 := &student{Name: "Geektutu", Age: 20}
	// stu2 := &student{Name: "Jack", Age: 22}
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "css.tmpl", nil)
	// })

	// r.GET("/students", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "arr.tmpl", gee.H{
	// 		"title":  "gee",
	// 		"stuArr": [2]*student{stu1, stu2},
	// 	})
	// })

	// r.GET("/date", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
	// 		"title": "gee",
	// 		"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	// 	})

	// })
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
