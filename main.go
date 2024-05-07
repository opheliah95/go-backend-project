package main
import (
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)
	

func main() {
	fmt.Println("hello world")

	// load go env
	err := godotenv.Load()
	if err !=nil {
		log.Fatal("Error loading .env file")
	}
	portString := os.Getenv("PORT")

	// check if port exists
	if portString == "" {
		log.Fatal("THE PORT VARIABLE IS NOT SET	")
	}

	//add server
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: ":"+ portString,
	}
	log.Printf("server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Port: ", portString)
}