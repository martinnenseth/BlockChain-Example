package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "usernames...")
	})

	m.Post("/", func(r *http.Request, x render.Render)  {
		text := string(r.FormValue("username"))
		x.HTML(200, "hello", "Brukernavnet " + text + " er lagt til i listen.")
	})

	m.Run()
}