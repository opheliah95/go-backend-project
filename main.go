package main
import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
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
	fmt.Println("Port: ", portString)
}