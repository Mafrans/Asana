package api

import (
	"fmt"
	"net/http"
)

func me() {
	response, err := http.Get(baseUrl)
	if (err != nil) {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)
}