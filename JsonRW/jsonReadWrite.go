package JsonRW

import (
	"encoding/json"

	"fmt"
	"os"
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


