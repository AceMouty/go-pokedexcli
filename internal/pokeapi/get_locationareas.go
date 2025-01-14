package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetLocationAreas(requestUrl *string) (RespLocationAreas, error) {

	if val, ok := c.cache.Get(*requestUrl); ok {
		fmt.Println("Using cache..")
		locationsResp := RespLocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocationAreas{}, err
		}

		return locationsResp, nil
	}

	time.Sleep(2 * time.Second)
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

	c.cache.Add(*requestUrl, jsonData)
	return locationsResp, nil
}
