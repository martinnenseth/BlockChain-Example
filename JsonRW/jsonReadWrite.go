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
	//appends the parameters to the struct.
	row := jsonItems{Name : name, Ip: ip}

	b, err := json.Marshal(row)
	if err != nil{
		fmt.Println(err)
	}

	ioutil.WriteFile("output1.json", b, 0644)
	fmt.Println(b)
}


