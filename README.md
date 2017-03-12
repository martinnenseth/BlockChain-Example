# Distrubuted List

## About
A disributed list application implemented in GoLang. 
<br><br>
<b> How it works </b> <br>
The application is based around a file named output1.json. This file include all of the username(s),
we've collected.. This will increase when someone add their username in the web interface.
<br><br> Each time a user add a username, it will add it to our file(output1.json) and send a request to
other hosts in our network. The network contains a list of IPs that have at least submitted 1 username. 
 
<p> There is also a routine that requests update from the other host every 5 minutes. This goes in a 
sepeate thread while the application runs. </p>


```golang
func runUpdateEveryFiveMinute(){
	time.Sleep(20 * time.Second) // to skip update request while the web-server boots up.
	for true {
		SendUpdateRequests()
		time.Sleep(5 * time.Minute)
	}
}
```
For retriving a file latest modifiy date, we used file.Stat()'s ModTime to get out the date..

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
.. So the modify date can be collected from remote, we needed a API.

```golang
m.Get("/api/data/fileLastEdited", func(r render.Render) {
	r.HTML(200, "apiUsernames", getLastEditTime())
})
```
.. And since we add it as a string, we need to parse it back to time format, in order to check if is a newer or older file.

```golang
file_date_remote, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", string(api_read))
```

We can now use this data to calculate if we need to change the file.. or just send the request back..
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
1. A client node inserts a username (and IP, but this is automated, 
<it> see line 42 in server.go </it>) into a json file on a host node. 
<br>
2. The host node requests all the nodes in the json file to update their file, 
to match the newest version of the file. 
<br>
2.1. All nodes fullfil the hosts request if the last edited time of the host file
 is newer than the client's.
<br>
![alt tag](https://scontent-arn2-1.xx.fbcdn.net/v/t35.0-12/17195398_10211771329536156_738295374_o.png?oh=b7dfb1e7bcdeee6813cf897b273b53bf&oe=58C50EE5)
