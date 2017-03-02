package JsonRW

import (
	"encoding/json"

	"fmt"
	"os"
	"io/ioutil"
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
	//appends the parameters to the struct.



	row := jsonItems{Name : name, Ip: ip}

	b, err := json.Marshal(row)
	if err != nil{
		fmt.Println(err)
	}

	f, err := os.OpenFile("output1.json", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Write(b); err != nil {
		panic(err)
	}
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

