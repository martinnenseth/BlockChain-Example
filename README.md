# Distrubuted List
<b>The following user / system scenarios is implemented:</b><br>
<br>
1. A user(client) access a Node trough an URL: <br>
http://ip-address:port/ and gets an html file with a
form element and fills in his / her username and sumbits the data to the
server. When the user does this, the username is added to the list.

2. A user (Client) views the last version of an updated list with user names
registered in the distributed system (this is a GET request with path
http://ip-address:port/members

3.  a server (acting like a client) requests a last updated list from another
server which responds with a JSON-respons
http://ip-address:port/api/data/json



<p>To ensure that a range of servers have the same content(in our example we've picked a list of usernames), we need to create some software that look for difference and apply changes. We've done so by having threads in our software to checks other servers in our network using the API. This check the date of the file, if the file is newer than the requested host, we simply change the content. 
</p><p>
When a new username is submitted to one of the servers webpage, it automaticlly starts requesting other servers to get the new content it've gathered. Otherwise, as mention it'll check for it every 5 minutes. 
</p> <p>
For more detailed description and code example.. Please look below! 
</p>
<img src ="http://i.imgur.com/CS9cUnw.png">


## About
A disributed list application implemented in GoLang. 
<br><br><p>
The application is based on a file named output1.json. This file includes all of the usernames and IPs related to them (implicitly gathered by the application) that we've collected. The number of these instances (usernames and IPs) will increase when someone adds their username in the web interface. 
</p><p>
Each time a user add a username, it will append it into our file(output1.json), and send a request to other hosts in our network(IPs in the json file).
</p><p>
There is also a routine that requests an update from the other host every 5 minutes. This goes in a separate thread while the application runs.
</p>

```golang
func runUpdateEveryFiveMinute(){
	time.Sleep(20 * time.Second) // to skip update request while the web-server boots up.
	for true {
		SendUpdateRequests()
		time.Sleep(5 * time.Minute)
	}
}
```
For receiving a file latest modify date, we used file.Stat()'s ModTime to get out the date.

```golang
func getLastEditTime() time.Time {
	file, err := os.Open("output1.json")
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	// Return the time of when the file was last modified.
	return fi.ModTime()
}
```
.. So the modify date can be collected from remote, we needed an API.

```golang
m.Get("/api/data/fileLastEdited", func(r render.Render) {
	r.HTML(200, "apiUsernames", getLastEditTime())
})
```
.. And since we add it as a string, we need to parse it back to time format, in order to check if is a newer or older file.

```golang
file_date_remote, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", string(api_read))
```

We can now use this data to calculate if we need to change the file.. or just send the request back.
```golang
if file_date_remote.Before(getLastEditTime()){
	// our file is newer.. send the request back.
	go SendUpdateRequests()
	return "Request sent back. reason: newer file spotted."
}else if file_date_remote.Equal(getLastEditTime()) {
	// file is the same
	return "Request denied. File date the same.."
}

println("Old file spotted, changing the file..")
```

## API's

```
~/api/data/filesize
```
Will return size in bytes..
```
~/api/data/json
```
Will return full json file of the content of our username list
```
~/api/data/accountName
```
Will count all of the usernames and display the amount of usernames in our json file.
```
~/api/data/fileLastEdited
```
Will display the full date of the jsonfile lastest modify date

/register	POST name

/list		GET json - usernames = { "uers: {"name": "Ola"}


