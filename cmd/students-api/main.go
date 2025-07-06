package main

import "fmt"
import "net/http"
import "log"
 
import "github.com/sharath666/students-api/cmd/students-api/internal/config"

func main(){
	//load config
	cfg := config.MustLoad()

	fmt.Printf("printing the cfg %s", cfg.Addr)
	//database setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("welcome to students api"))
	})
	//setup server

	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server %v", err)

	}

	fmt.Println("server started")
}