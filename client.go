package bandiera

import (
	"encoding/json"
)

type Client struct {
	httpClient HttpClient
}

func NewBandieraClient(url string) *Client {
	return &Client{newNetHttpClient(url)}
}

func (c Client) GetAll(params Params) (res AllResponse) {
	body, err := c.httpClient.GetUrlContent("/api/v2/all", params)
	if err != nil {
		return
	}

	_ = json.Unmarshal([]byte(body), &res)
	return
}

func (c Client) GetFeaturesForGroup(group string, params Params) (res GroupResponse) {
	body, err := c.httpClient.GetUrlContent("/api/v2/groups/"+group+"/features", params)
	if err != nil {
		return
	}

	_ = json.Unmarshal([]byte(body), &res)
	return
}

func (c Client) IsEnabled(group, feature string, params Params) (res bool) {
	body, err := c.httpClient.GetUrlContent("/api/v2/groups/"+group+"/features/"+feature, params)
	if err != nil {
		return
	}

	payload := FeatureResponse{}
	err = json.Unmarshal([]byte(body), &payload)
	if err != nil {
		return false
	}

	return payload.Enabled
}
