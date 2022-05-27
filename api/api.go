package api

import (
	"errors"
	"net/http"
	"os"
)

const baseUrl string = "https://app.asana.com/api/1.0"

func apiGet(path string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", baseUrl + path, nil)
	token, err := authorize()
	if (err != nil) {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer " + token)
	
	return http.DefaultClient.Do(req);
}

func authorize() (string, error) {
	if (os.Getenv("PAT") != "") {
		return os.Getenv("PAT"), nil
	}
	return "", errors.New("OAuth authentication not implemented, please use a PAT instead. (https://developers.asana.com/docs/personal-access-token)")
}


type routes struct {}
var Routes = routes{}