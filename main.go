package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/httpReq/handlers"
)

/*
	Write a Golang server.
	- Which accepts GET and POST http methods
	- Post Method should read the request body and store it  in a struct(you can create a struct of your choice).
	- GET should return a struct converted to json (you can hardcode the data for struct).

	[optional]
	- You can store the POST api’s data in a file and return all the data when GET api is called.
*/

func main() {

	logg := log.New(os.Stdout, "HTTP-Assign", log.LstdFlags)

	// Create a custom ServeMux
	sm := http.NewServeMux()

	// Create a new Handler
	uh := handlers.NewUserReq(logg)

	// Adding the Handler to the Custom serveMux
	sm.Handle("/",uh)

	newServer := http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 20*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	newServer.ListenAndServe()
}