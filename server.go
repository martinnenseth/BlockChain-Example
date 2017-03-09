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
	"bytes"
)

func main() {

	m := martini.Classic()
	// render html templates from templates directory

	m.Use(render.Renderer(render.Options{
		IndentJSON: true, // Output human readable JSON
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "")

	})
	m.Post("/", func(r *http.Request, x render.Render)  {
		text := string(r.FormValue("username"))
		readApi, _ := http.Get("https://api.ipify.org")
		bytes, _ := ioutil.ReadAll(readApi.Body)

		JsonRW.WriteInstance(text, string(bytes))

		x.HTML(200, "hello", "" + text + " is added to the list.")
	})

	m.Get("/members", func(r render.Render) {
		r.HTML(200, "header", "")

		// for each member in our json file
		index := 1
		for _, member := range JsonRW.ReadEntireJson() {

			nr := string(index)

			r.HTML(200, "main", nr + "# " + member["name"] + " - IP:  " + member["ip"])
			index = index + 1;
		}

		r.HTML(200, "footer", "")
	})

	m.Get("/servers", func(r render.Render){
		r.HTML(200, "header", "")
		r.HTML(200, "main", "404: It's not the size of the guy that matters, it's the loyalty of his guns..")
		r.HTML(200, "footer", "")
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

	m.Post("/api/runUpdate", func(r *http.Request, x *bytes.Buffer) string {
		/*
			Some other host requested this host to update
		*/

		// get parameters
		fromHost := string(r.FormValue("addr"))		// the ip that sent the request
		token := string(r.FormValue("token"))		// token for requesting update

		if (token != "someTokenToPreventUnauthoriseUpdateRequest") {
			// not authorised for requesting update
			return "Token not valid"

		}

		// check if requesting host have a bigger file
		hostFileSize, err :=  http.Get(fromHost + "/api/data/filesize")
		if err != nil {log.Fatal(err)}

		bytes, err := ioutil.ReadAll(hostFileSize.Body)
		if err != nil {log.Fatal(err)}

		i, err := strconv.ParseInt(string(bytes), 10, 64)
		if err != nil {log.Fatal(err)}

		if i < GetCurrentFileSize() {
			// if current file size is higher.. do nothing, and request the host to update their file.
			// TO DO, REQUEST HOST.
			return "our data is newer, i'll send your request back."

		} else if i == GetCurrentFileSize() {
			return "data is the same"
		}

		// okay, we got a file with less data than the other host.. we gotta grab that instead.
		readAPi, err := http.Get(fromHost + "/api/data/json")
		if err != nil {log.Fatal(err)}

		jsonByte, err := ioutil.ReadAll(readAPi.Body)
		if err != nil {log.Fatal(err)}

		// write to file
		ioutil.WriteFile("output1.json", jsonByte, 0644)

		return "File changed"

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






