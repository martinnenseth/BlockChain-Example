package JsonRW

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
)
type jsonItems struct {
	Name string
	Ip string
}

/**
 * @param name and ip to append.
 * Creates a new instance to the json file.
 */
func WriteInstance(name string, ip string)  {
	//assign the parameters to the struct.
	row := jsonItems{Name : name, Ip: ip}

	b, err := json.Marshal(row)
	if err != nil{
		fmt.Println(err)
	}

	ioutil.WriteFile("output1.json", b, 0644)
	fmt.Println(b)
}
/*
 * Decodes the json file and prints its content.
 */
func ReadEntireJson ()  {
	jsonFile, err := ioutil.ReadFile("output1.json")
	if err != nil{
		fmt.Print(err)
	}
	var items jsonItems // Variable to store the content of the file.
	err = json.Unmarshal(jsonFile, &items)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(items)
}








