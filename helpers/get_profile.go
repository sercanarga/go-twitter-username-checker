package helpers

import (
	"errors"
	"io"
	"net/http"
)

type ProfileStruct struct {
	APIUrl      string
	ScreenName  string
	BearerToken string
}

func GetProfile(profile ProfileStruct) (int, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", profile.APIUrl, nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("screen_name", profile.ScreenName)
	req.URL.RawQuery = q.Encode()

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + profile.BearerToken},
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != 200 && res.StatusCode != 404 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return 0, err
		}
		defer res.Body.Close()

		return 0, errors.New(string(body))
	}

	return res.StatusCode, nil
}
