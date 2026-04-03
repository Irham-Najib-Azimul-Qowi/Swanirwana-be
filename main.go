package main

// import "fmt"
import (
	"tiket-app-backend/config"
    "tiket-app-backend/routes"
    "github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Swanirwana Backend Project Initialization...")
	// 1. koneksi database 
	config.ConnectDB()

	// 2. setup gin router 
	router := git.Default()

	// 3. setup routes 
	routes.SetupRoutes(router)

	// 4. jalankan server di port 8080 
	router.run(':8080')
}
