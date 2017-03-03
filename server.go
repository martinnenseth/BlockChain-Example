package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"./JsonRW"

	"io/ioutil"
	"os"
	"log"
	"fmt"
)

func main() {

	m := martini.Classic()
	// render html templates from templates directory




	m.Use(render.Renderer(render.Options{
		IndentJSON: false, // Output human readable JSON

	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "")

	})


	m.Get("/members", func(r render.Render) {
		r.HTML(200, "header", "")
		for _, member := range JsonRW.ReadEntireJson() {
			r.HTML(200, "main", member["name"] + " - IP:  " + member["ip"])
		}

		r.HTML(200, "footer", "")
	})

		// https://api.ipify.org
	m.Post("/", func(r *http.Request, x render.Render)  {
		text := string(r.FormValue("username"))
		readApi, _ := http.Get("https://api.ipify.org")
		bytes, _ := ioutil.ReadAll(readApi.Body)

		JsonRW.WriteInstance(text, string(bytes))

		x.HTML(200, "hello", "" + text + " is added to the list.")
	})

	m.Get("/api/member/filesize", func() string {
		return fmt.Sprintf("%d", GetCurrentFileSize())
	})

	m.Get("/api/member/json", func(r render.Render) {
		fmt.Println(JsonRW.GetRawJsonFile())
		r.HTML(400, "apiUsernames", JsonRW.GetRawJsonFile())
	})

	m.Get("/api/member/amountName", func(r render.Render) {
		fmt.Println(JsonRW.GetAmountOfUsername())
		r.HTML(400, "apiUsernames", JsonRW.GetAmountOfUsername())
	})

	m.RunOnAddr(":8080")
	m.Run()
}

func GetCurrentFileSize() int64 {
	file, err := os.Open("output1.json")

	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	return fi.Size()
}






