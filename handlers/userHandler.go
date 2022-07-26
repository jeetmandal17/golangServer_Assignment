package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/httpReq/data"
)

type UserReq struct{
	l		*log.Logger
}

func NewUserReq(l *log.Logger) *UserReq{
	return &UserReq{l}
}

func (u UserReq) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	// Handling POST Request
	if r.Method == http.MethodPost{
		u.l.Println(" POST ")

		u.Add(rw,r)
		return
	}

	// Handling GET Request
	if r.Method == http.MethodGet{
		u.l.Println(" GET ")

		u.Get(rw,r)
		return
	}

	// Handling Exceptions
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Handling POST requests
func (u UserReq) Add(rw http.ResponseWriter, r *http.Request){

	u.l.Println(" Handling POST Request ")

	// Convert JSON into User object
	newUser := &data.User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	u.l.Println(newUser)

	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusInternalServerError)
		return
	}

	data.AddUser(*newUser)
	rw.Write([]byte("Succesfully Added user in the file"))
}

// Handling GET requests
func (u UserReq) Get(rw http.ResponseWriter, r *http.Request){

	u.l.Println(" Handling GET Request ")

	// GET the Users from the JSON file
	alldata, err := data.GetUser()
	if err == data.ErrorOpeningFileForRead {
		u.l.Println(data.ErrorOpeningFile)
		return
	}

	for _,item := range alldata{
		json.NewEncoder(rw).Encode(item)
	}
}