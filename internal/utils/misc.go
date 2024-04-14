package utils

import (
	"encoding/json"
	"net/http"
)

func GetJson(res *http.Response, target interface{}) error {
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
