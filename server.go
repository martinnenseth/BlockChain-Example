package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"fmt"
)

func main() {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "usernames...")
	})

	m.Post("/submitUsername", func(req *http.Request){
		fmt.Println(req.PostForm)
	})

	m.Run()


}