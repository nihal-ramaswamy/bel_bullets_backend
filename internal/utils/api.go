package utils

import (
	"bel_bullets/internal/dto"
	"encoding/json"
	"net/http"
)

func GetJson(res *http.Response, target interface{}) error {
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func GetClubById(clubId string, authToken string) (dto.Club, error) {

	clubDetails := dto.Club{}

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/clubs/"+clubId, nil)
	if nil != err {
		return clubDetails, err
	}
	req.Header.Add("Authorization", authToken)

	res, err := client.Do(req)
	if nil != err {
		return clubDetails, err
	}

	err = GetJson(res, &clubDetails)

	return clubDetails, err
}
