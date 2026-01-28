package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type misskeyCreateNoteRequest struct {
	I    string `json:"i"`
	Text string `json:"text"`
}

func PostToMisskey(note string) error {
	domain := strings.TrimRight(conf.Credentials.MisskeyCredentials.MisskeyUrl, "/")
	if domain == "" {
		return errors.New("misskey domain is empty")
	}
	if conf.Credentials.MisskeyCredentials.MisskeyToken == "" {
		return errors.New("misskey access token is empty")
	}

	payload, err := json.Marshal(misskeyCreateNoteRequest{
		I:    conf.Credentials.MisskeyCredentials.MisskeyToken,
		Text: note,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, domain+"/api/notes/create", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return errors.New("failed to retrieve error")
		}
		return errors.New(string(body))
	}

	return nil
}
