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
	"net/url"
	"strings"
)


func main() {

	m := martini.Classic()


	m.Use(render.Renderer(render.Options{
		IndentJSON: true, // Output human readable JSON
	}))


	m.NotFound(func(r render.Render) {
		r.HTML(200, "header", "")
		r.HTML(200, "header-text", "404...")
		r.HTML(200, "main", "Siden eksisterer ikke.")
		r.HTML(200, "footer", "")
	})

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
		JsonRW.WriteInstance(text, getServerIP())
		SendUpdateRequests()
		x.HTML(200, "hello", "" + text + " is added to the list.")
	})

	/*
		Render all the usernames we have collected so far.
	 */
	m.Get("/members", func(r render.Render) {
		// To force check that the node got the latest file, send update request
		SendUpdateRequests()

		r.HTML(200, "header", "")
		r.HTML(200, "header-text", "Members we have collected so far")
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
		r.HTML(200, "header-text", "Servers that have added usernames")
		for _, ip := range JsonRW.GetAllIPs() {
			r.HTML(200, "main", ip)
		}
		r.HTML(200, "footer", "")
	})

	/**
		Gets a list of api addr.
	 */
	m.Get("/api", func (r render.Render) {
		r.HTML(200, "header", "")
		r.HTML(200, "header-text", "Lists of all api's")
		for _, x := range m.All() {
			if strings.HasPrefix(x.Pattern(), "/api/data") {
				r.HTML(200, "links", x.Pattern())
			}
		}
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
	m.Post("/api/runUpdate", func(r *http.Request) string {

		println("Update request recived")

		// get the IP from the requester
		fromHost := string(r.FormValue("addr"))
		// get token provided
		token := string(r.FormValue("token"))
		
		// Check if host is authorised to update our data.
		if token != "someTokenToPreventUnauthoriseUpdateRequest" {return "bad token"}

		println("token accepted")

		// check if requesting host have a bigger file
		hostFileSize, err :=  http.Get(fromHost + "/api/data/filesize")
		println(fromHost + "/api/data/filesize")

		if err != nil {
			println("error")
			log.Fatal(err)
		}
		println("Get's requesters filesize")
		// parse hostFileSize over to int...
		byte_host_file_size, err := ioutil.ReadAll(hostFileSize.Body)
		println("parsing hostfile size to int")
		int_host_file_size, err := strconv.ParseInt(string(byte_host_file_size), 10, 64)
		if err != nil {log.Fatal(err)}

		println("Host file size converted to int, and now beeing compared")
		// if the current file size is larger, we do not wanna do anything..
		// .. instead we send the request back to the requesting host.
		if int_host_file_size < GetCurrentFileSize() {
			SendUpdateRequests()
			return "our data is newer, i'll send your request back."

		} else if int_host_file_size == GetCurrentFileSize() {
			println("data is the same")
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

/*
	For sending out update requests to other hosts in the network.
 */
func SendUpdateRequests() {

	println("Sending update request to other nodes..")

	// collect list of servers, based on the json file with usernames..
	servers := JsonRW.GetAllIPs()

	// get our ip..
	host_ip := getServerIP()

	// for each server in our server list
	hc := http.Client{Timeout: 20}
	form := url.Values{}
	form.Add("addr", "http://"+host_ip+":8080")
	form.Add("token", "someTokenToPreventUnauthoriseUpdateRequest")
	for _, ip := range servers {
		if ip != host_ip {
			println("Sending update request for " + ip)
			url := "http://" + ip +":8080" + "/api/runUpdate"
			req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))

			if err != nil {
				println(err)
			}
			req.PostForm = form
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err := hc.Do(req)

			print(resp)

			if err != nil {
				println(err)
			}
			//resp_string, err := ioutil.ReadAll(resp.Body)
			//println(string(resp_string))

		}
	}
}

/*
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

func getServerIP() string{
	readApi, err := http.Get("https://api.ipify.org")
	if err != nil {log.Fatal(err)}
	bytes, err := ioutil.ReadAll(readApi.Body)
	if err != nil {log.Fatal(err)}
	return string(bytes)
}






