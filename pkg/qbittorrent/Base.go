package qbittorrent

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	port      string
	V2BaseUrl string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port = os.Getenv("QBIT_PORT")
	V2BaseUrl = fmt.Sprintf("http://localhost:%s/api/v2/", port)
}
