package api

import (
	"encoding/json"
	"io/ioutil"
)

type MeResponse struct {
	Data struct {
		Gid string `json:"gid"`
		Email string `json:"email"`
		Name string `json:"name"`
		Photo map[string]string `json:"photo"`
		ResourceType string `json:"resource_type"`
		Workspaces []struct {
			Gid string `json:"gid"`
			Name string `json:"name"`
			ResourceType string `json:"resource_type"`
		} `json:"workspaces"`
	} `json:"data"`
}

func (_ *routes) Me() (MeResponse) {
	res, err := apiGet("/users/me")
	if (err != nil) {
		panic(err)
	}

	if (res.StatusCode == 401) {
		panic("Unauthorized, make sure to check whether your PAT is valid or not") 
	}

	text, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if (err != nil) {
		panic(err)
	}

	var meResponse MeResponse
	json.Unmarshal(text, &meResponse)

	return meResponse
}