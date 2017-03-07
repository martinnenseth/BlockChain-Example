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
	"strconv"
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


	// api's
	m.Get("/api/data/filesize", func() string {
		return fmt.Sprintf("%d", GetCurrentFileSize())
	})

	m.Get("/api/data/json", func(r render.Render) {
		fmt.Println(JsonRW.GetRawJsonFile())
		r.HTML(200, "apiUsernames", JsonRW.GetRawJsonFile())
	})

	m.Get("/api/data/amountName", func(r render.Render) {
		fmt.Println(JsonRW.GetAmountOfUsername())
		r.HTML(200, "apiUsernames", JsonRW.GetAmountOfUsername())
	})

	m.Post("/api/runUpdate", func(r *http.Request) {
		/*
			Some other host requested this host to update
		 */
		fromHost := string(r.FormValue("addr"))			// the requested hosts addr

		// Check if host is authorised to update our data.
		token := string(r.FormValue("token"))
		if (token != "someTokenToPreventUnauthoriseUpdateRequest") {
			return
		}

		// check if requesting host have a bigger file
		hostFileSize, err :=  http.Get(fromHost + "/api/data/filesize")

		if err == nil {
			//error
			return
		}
		i, err := strconv.ParseInt(hostFileSize, 10, 64)
		if err != nil {
			panic(err)
		}

		if (i < GetCurrentFileSize()) {
			// if current file size is higher.. do nothing, and request the host to update their file.
			// TO DO, REQUEST HOST.
			return
		}

		// okay, server got a file with less data than the other host.. we gotta grab that instead.
		readAPi, err := http.Get(fromHost + "/api/data/json")
		if err == nil {
			//error
			return
		}
		bytes, err := ioutil.ReadAll(readAPi)
		if err == nil {
			//error
			return
		}
		// write to file
		ioutil.WriteFile("output1.json", bytes, 0644)
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






