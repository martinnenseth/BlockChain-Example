# Distrubuted List

## About
A disributed list application implemented in GoLang. 
<b> How it works </b> <br>
1. A client node inserts a username (and IP, but this is automated, <it> see line 42 in server.go </it>) into a json file on a host node. 
2. The host node requests all the nodes in the json file to update their file, to match the newest version of the file. 
2.1. All nodes fullfil the hosts request if the last edited time of the host file is newer than the client's.
<br>
![alt tag](https://scontent-arn2-1.xx.fbcdn.net/v/t35.0-12/17195398_10211771329536156_738295374_o.png?oh=b7dfb1e7bcdeee6813cf897b273b53bf&oe=58C50EE5)
