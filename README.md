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
<p> With help from golang, we retrived last edited date from the file.. </p>
```golang
/**
	Get the latest edited time of the file.
	@return time of last edit.
 */
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
<p> And with a coded API, we can retrive a remote servers's files modified date, and by that we can use that data to find out if a servers file is outdated or not. 
</p>
```golang
m.Get("/api/data/fileLastEdited", func(r render.Render) {
		r.HTML(200, "apiUsernames", getLastEditTime())
	})
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
