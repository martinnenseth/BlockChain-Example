package JsonRW

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"fmt"
	"io"
	"os"
)
type members []map[string]string

func WriteInstance(name string, ip string) {
	var data members

	jsonFile, err := ioutil.ReadFile("output1.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Fatal(err)
	}
	instance := map[string]string{
		"name": name,
		"ip": ip,
	}
	data = append(data, instance)
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("output1.json", b, 0644)
}

/*
 * Decodes the json file and prints its content.
 */
func ReadEntireJson ()  members{
	jsonFile, _ := os.Open("output1.json")

	var u members

	dec := json.NewDecoder(jsonFile)
	for {


		if err:= dec.Decode(&u); err == io.EOF {
			break
		}else if err != nil {
			log.Fatal(err)
		}

	}
	for member := range u {
		fmt.Printf("name: %s", u[member]["name"])

	}
	return u
}



