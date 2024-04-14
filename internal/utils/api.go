package utils

import (
	constants "bel_bullets/internal/constants/club"
	"bel_bullets/internal/dto"
	"net/http"
)

func GetClubById(clubId string, authToken string, page string) (dto.Club, error) {

	clubDetails := dto.Club{}

	client := &http.Client{}

	req, err := http.NewRequest("GET", constants.BASE+constants.CLUB+clubId, nil)
	if nil != err {
		return clubDetails, err
	}

	req.Header.Add("Authorization", authToken)
	query := req.URL.Query()
	query.Add("page", page)
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if nil != err {
		return clubDetails, err
	}

	err = GetJson(res, &clubDetails)

	return clubDetails, err
}

func GetClubMembers(clubId string, authToken string, page string) ([]dto.ClubMembers, error) {

	clubMembers := []dto.ClubMembers{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", constants.BASE+constants.CLUB+clubId+"/"+constants.MEMBERS, nil)
	if nil != err {
		return clubMembers, err
	}

	req.Header.Add("Authorization", authToken)
	query := req.URL.Query()
	query.Add("page", page)
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if nil != err {
		return clubMembers, err
	}

	err = GetJson(res, &clubMembers)

	return clubMembers, err

}
