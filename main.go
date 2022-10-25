package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"popflash/api/models"

	"github.com/davecgh/go-spew/spew"
)

var baseUrl string = "https://api.popflash.site/api/rest"

func main() {

	// c, err := GetClanMatches("fug-hub", 2)
	// c, err := GetUserMatches(1, 1)
	c, err := GetMatch(1360060)

	if err != nil {
		fmt.Println(err)
	}

	spew.Dump(c)

}

func GetClanMatches(clan string, limit int8) (*models.ClanMatchesResponse, error) {
	url := fmt.Sprintf("%s/clan/%s/matches?limit=%v", baseUrl, clan, limit)

	res, err := request[models.ClanMatchesResponse](url)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUserMatches(user int32, limit int8) (*models.UserMatchesResponse, error) {
	url := fmt.Sprintf("%s/user/%v/matches?limit=%v", baseUrl, user, limit)

	res, err := request[models.UserMatchesResponse](url)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetMatch(id int32) (*models.MatchResponse, error) {
	url := fmt.Sprintf("%s/match/%v", baseUrl, id)

	res, err := request[models.MatchResponse](url)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func request[T any](url string) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := send(req)
	if err != nil {
		return nil, err
	}
	var data T
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func send(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
