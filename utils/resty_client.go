package utils

import "github.com/go-resty/resty/v2"

// DoRestyGet performs GET request
func DoRestyGet(url string) (string, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
