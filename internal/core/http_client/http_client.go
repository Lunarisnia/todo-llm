package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var Do *HTTPClient = &HTTPClient{
	client: &http.Client{},
}

type HTTPClient struct {
	client *http.Client
}

func (h *HTTPClient) Post(url string, body []byte, dest any) error {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	resp, err := h.client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return err
	}
	defer resp.Body.Close()

	// FIXME: handle the case if dest is not a reference to something
	err = json.NewDecoder(resp.Body).Decode(dest)
	if err != nil {
		return err
	}

	return nil
}
