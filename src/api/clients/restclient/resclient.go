package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, headers http.Header, body interface{}) (*http.Response, error)  {
	// marshal body return if err
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}