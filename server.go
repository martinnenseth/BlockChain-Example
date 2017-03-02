package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"./JsonRW"

	"io/ioutil"
)

func main() {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "")
	})

	m.Get("/members", func(r render.Render) {
		r.HTML(200, "members", "her kommer medlemmer..")
	})

		// https://api.ipify.org
	m.Post("/", func(r *http.Request, x render.Render)  {
		text := string(r.FormValue("username"))
		readApi, _ := http.Get("https://api.ipify.org")
		bytes, _ := ioutil.ReadAll(readApi.Body)

		JsonRW.WriteInstance(text, string(bytes))
		x.HTML(200, "hello", "" + text + " is added to the list.")
	})

	m.Run()
}


