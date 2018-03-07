package utils

import (
	"fmt"
	"net/http"
)

func getmyip() {
	client := &http.Client{}
	fmt.Print(HttpGet("https://api.ipify.org?format=json", client))
	fmt.Printf("\n")
}
