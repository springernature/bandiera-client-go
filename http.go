package bandiera

import (
	"io/ioutil"
	"net/http"
)

type NetHttpClient struct {
	baseUrl string
}

func newNetHttpClient(baseUrl string) HttpClient {
	return &NetHttpClient{baseUrl}
}

func (c *NetHttpClient) GetUrlContent(url string, params Params) (body []byte, err error) {
	fullUrl := c.baseUrl + url

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return []byte{}, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return
}
