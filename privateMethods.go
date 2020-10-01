package qiwi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// NewRequest - create new request
func (qiwi *PersonalAPI) newRequest(apiKey, method, spath string, data map[string]interface{}) (req *http.Request, err error) {

	var path = APIURL + spath

	var body io.Reader

	if len(data) > 0 {

		s, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(s)

	}

	req, err = http.NewRequest(method, path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return req, err
}

// NewPersonalAPIWithHTTPClient - new api with client
func newPersonalAPIWithHTTPClient(apiKey string, client *http.Client) *PersonalAPI {
	qiwi := new(PersonalAPI)

	qiwi.httpClient = client
	qiwi.apiKey = apiKey

	return qiwi
}

// sendRequest - send request to qiwi api
func (qiwi *PersonalAPI) sendRequest(apiKey, method, spath string, data map[string]interface{}) (body []byte, err error) {
	req, err := qiwi.newRequest(apiKey, method, spath, data)

	response, err := qiwi.httpClient.Do(req)

	if err != nil && response == nil {

		return nil, err

	}

	body, err = ioutil.ReadAll(response.Body)

	if err != nil {

		if response.Body != nil {
			response.Body.Close()
		}

		return nil, err
	}

	if response.StatusCode == 400 {

		var res TransferError

		err = json.Unmarshal(body, &res)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s : %s", res.Code, res.Message)
	}

	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	if response.Body != nil {
		response.Body.Close()
	}

	return body, nil

}
