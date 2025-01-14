package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetLocationArea(requestUrl *string) (Location, error) {

	if val, ok := c.cache.Get(*requestUrl); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	time.Sleep(2 * time.Second)
	req, err := http.NewRequest("GET", *requestUrl, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(*requestUrl, dat)

	return locationResp, nil
}
