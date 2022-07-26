package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Creating Custom Error [Write Errors]
var ErrorOpeningFile = fmt.Errorf("cannot open file to add the User")
var ErrorMarshalUser = fmt.Errorf("cannot convert ot JSON object")
var ErrorWritingIntoFile = fmt.Errorf("cannot write JSON object into file")

// Creating Custom Error [Read Errors]
var ErrorOpeningFileForRead = fmt.Errorf("cannot open file for reading data")
var ErrorCannotUnmarshal = fmt.Errorf("cannot unmarhsal JSON to user struct")


type User struct{
	Name 	string		`json:"name"`
	Age		int			`json:"age"`
	ID		string		`json:"id"`
	Salary	int			`json:"salary"`
}

func AddUser(newUser User) error{

	// Get all the users in the JSON file
	oldUsers, _ := GetUser()
	oldUsers = append(oldUsers, newUser)

	// OPEN the file for adding new User
	file, err := os.OpenFile("store.json", os.O_CREATE | os.O_RDWR | os.O_TRUNC, 0666)
	if err != nil {
		return ErrorOpeningFile
	}
	defer file.Close()

	// Convert the struct into JSON object
	// jsonUser, err := json.Marshal(newUser)
	// if err != nil {
	// 	return ErrorMarshalUser
	// }

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(oldUsers)
	

	//Writing User into the file
	_, err = file.Write(reqBodyBytes.Bytes())
	if err != nil {
		return ErrorWritingIntoFile
	}

	return nil
}

func GetUser() ([]User, error){

	userData := []User{}

	// OPEN the file as readOnly
	allData, err := ioutil.ReadFile("store.json")
	if err != nil{
		return userData, ErrorOpeningFileForRead
	}

	if len(allData) == 0{
		return userData, nil
	}

	err = json.Unmarshal([]byte(allData), &userData)

	if err != nil{
		return userData, ErrorCannotUnmarshal
	}
	return userData, nil
}
