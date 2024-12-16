package main

import (
	"github.com/fuadroidd/chitchat.be/internal/app/http"
)

func main() {
	router := http.SetupRouter()
	router.Run(":3000")
}
