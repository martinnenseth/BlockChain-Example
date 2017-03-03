package JsonRW

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"fmt"
	"io"
	"os"
)
type Members []map[string]string

func WriteInstance(name string, ip string) {
	var data Members

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
func ReadEntireJson ()  Members{
	jsonFile, _ := os.Open("output1.json")

	var u Members

	dec := json.NewDecoder(jsonFile)
	for {


		if err:= dec.Decode(&u); err == io.EOF {
			break
		}else if err != nil {
			log.Fatal(err)
		}

	}
	for _, member := range u{

		fmt.Printf("name: %s", member["name"])

	}
	return u
}


func GetRawJsonFile() string {
	filename := "output1.json"
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return string(b)


}


