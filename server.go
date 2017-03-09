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

	m.Use(render.Renderer(render.Options{
		IndentJSON: true, // Output human readable JSON
	}))

	/*
		Our index page.
	 */
	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", "")

	})

	/*
		A post that want to add a new username to our collection.
	 */
	m.Post("/", func(r *http.Request, x render.Render)  {
		text := string(r.FormValue("username"))
		readApi, _ := http.Get("https://api.ipify.org")
		bytes, _ := ioutil.ReadAll(readApi.Body)

		JsonRW.WriteInstance(text, string(bytes))

		x.HTML(200, "hello", "" + text + " is added to the list.")
	})

	/*
		Render all the usernames we have collected so far.
	 */
	m.Get("/members", func(r render.Render) {
		r.HTML(200, "header", "")

		// for each member in our json file
		for _, member := range JsonRW.ReadEntireJson() {
			r.HTML(200, "main", member["name"] + " - IP:  " + member["ip"])
		}
		r.HTML(200, "footer", "")
	})

	/*
		This will in time show all connected servers.
	 */
	m.Get("/servers", func(r render.Render){
		r.HTML(200, "header", "")
		r.HTML(200, "main", "404: It's not the size of the guy that matters, it's the loyalty of his guns..")
		r.HTML(200, "footer", "")
	})

	/*
		API's for file size, json raw file and all of the account names.
	 */
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


	/*
		***************** REQUEST UPDATE TO UPDATE LOCAL FILE ********************
		Other hosts use this for quality assurance it's content. #Blockchain love.
	*/
	m.Post("/api/runUpdate", func(r *http.Request, x *bytes.Buffer) string {

		// get the IP from the requester
		fromHost := string(r.FormValue("addr"))
		// get token provided
		token := string(r.FormValue("token"))
		
		// Check if host is authorised to update our data.
		if token != "someTokenToPreventUnauthoriseUpdateRequest" {return "bad token"}

		// check if requesting host have a bigger file
		hostFileSize, err :=  http.Get(fromHost + "/api/data/filesize")
		if err == nil {log.Fatal(err)}

		// parse hostFileSize over to int...
		byte_host_file_size, err := ioutil.ReadAll(hostFileSize.Body)
		int_host_file_size, err := strconv.ParseInt(string(byte_host_file_size), 10, 64)
		if err != nil {log.Fatal(err)}

		// if the current file size is larger, we do not wanna do anything..
		// .. instead we send the request back to the requesting host.
		if int_host_file_size < GetCurrentFileSize() {
			SendUpdateRequests()
			return "our data is newer, i'll send your request back."

		} else if int_host_file_size == GetCurrentFileSize() {

			return "data is the same"
		}

		// okay, we got a file with less data than the other host..
		// .. we grab that instead.
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

func SendUpdateRequests() {
	servers := JsonRW.GetAllIPs()

	readApi, err := http.Get("https://api.ipify.org")
	if err != nil {log.Fatal(err)}

	bytes, err := ioutil.ReadAll(readApi.Body)
	if err != nil {log.Fatal(err)}

	host_ip := string(bytes)

	for _, ip := range servers {
		if ip != host_ip {
			// send ip request here..
		}
	}
}

/**
	get current size of json file.
 */
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






