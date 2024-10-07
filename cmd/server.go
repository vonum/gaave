package main

import (
	"log"
  "os"

  "github.com/vonum/gaave/server"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  dbUrl := os.Getenv("DB_URL")

  s := server.Server{}
  err = s.Run(8080, dbUrl)

  if err != nil {
    log.Fatal("Failed to run server on port 8080")
  }
}
