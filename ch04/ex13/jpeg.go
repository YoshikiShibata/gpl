// Copyright © 2016, 2021 Yoshiki Shibata. All rights reserved.

package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func fetchJPEG(title, posterURL string) (filename string, err error) {
	resp, err := http.Get(posterURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	filename = strings.Replace(title, " ", "_", -1)
	filename += ".jpg"
	err = os.WriteFile(filename, body, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}
