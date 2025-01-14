package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(requestUrl *string) (RespLocationAreas, error) {

	req, err := http.NewRequest("GET", *requestUrl, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	locationsResp := RespLocationAreas{}
	err = json.Unmarshal(jsonData, &locationsResp)
	if err != nil {
		return RespLocationAreas{}, err
	}

	return locationsResp, nil
}
