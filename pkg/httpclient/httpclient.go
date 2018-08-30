package httpclient

import (
	"fmt"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Clair/3.0 (https://github.com/coreos/clair/)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response was %d", resp.StatusCode)
	}

	return resp, nil
}
