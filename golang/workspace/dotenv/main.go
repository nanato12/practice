package main

import (
	"os"
    "fmt"
    "log"
    "github.com/joho/godotenv"
)
func Env_load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    Env_load()
    message := fmt.Sprintf("user_name=%s password=%s", os.Getenv("USER_NAME"), os.Getenv("PASSWORD"))
    fmt.Println(message)
}
