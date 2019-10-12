package gomeli

import (
	"net/http"
)

func (client *Client) Get(path string) (interface{}, *http.Response, error) {
	var body interface{}

	params := client.Auth.getParams()
	req, err := client.sling.New().Get(path).QueryStruct(params).Request()
	if err != nil {
		return body, nil, err
	}

	resp, err := client.sling.Do(req, &body, nil)

	return body, resp, err
}
