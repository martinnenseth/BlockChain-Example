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
