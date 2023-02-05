package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
type Something struct {
	Name    string  `json:"name"`
	Married bool    `json:"married"`
	Age     float64 `json:"age"`
	Address []struct {
		Street int    `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
	Marks []interface{} `json:"marks"`
}
type Address struct{
	Street int 'json:"street"'
	city string 'json:"city"'
}

// func init(){
// 	database.Setup()
// }
// func init(){
// 	fmt.Print(1)
// }
// func init(){
// 	fmt.Print(2)
// }
func main() {
	jsonFile,err:=os.Open("something.json")
	if err != nil{
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err:= ioutil.ReadAll(jsonFile)
	if err != nil{
		log.Fatal(err)

	}
	
	var something Something 
	json.unmarshal(jsonByteValues, &something)//converting hsn to struct
	log.Println(something)
	// //fmt.Printf(string(jsonByteValues)) // converting json data to string

	// newJsonByteValues, err:=jsonMarshal(something)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// //os.WriterFile("some.json", newJsonByteValues)

/
	// database.Setup()                    //establish the dtatbase connection
	// engine := routers.Router()          //get the customized engine
	// err := engine.Run("127.0.0.1:8080") // start the engine
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//consuming rest api in go code
	//using the rest api
	//returned value ?= json
	//now you must know how to read the json data
	// what is the problem if you read json data in string format
	// read json in structure format - marshalling , unmarshalling
}
