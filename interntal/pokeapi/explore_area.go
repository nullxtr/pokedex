package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(area string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + area

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	areaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return areaResp, nil
}
