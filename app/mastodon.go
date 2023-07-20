package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

func PostToMastodon(toot string) error {
	val := url.Values{}
	val.Set("status", toot)

	res, err := http.PostForm(conf.Credentials.MastodonCredentials.MastodonUrl+"/api/v1/statuses?access_token="+conf.Credentials.MastodonCredentials.MastodonToken, val)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		_, err := io.ReadAll(res.Body)
		if err != nil {
			return errors.New("failed to retrieve error")
		}
		return nil
	}

	return nil
}
