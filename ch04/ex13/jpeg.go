// Copyright © 2016 Yoshiki Shibata

package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func fetchJPEG(title, posterURL string) (filename string, err error) {
	resp, err := http.Get(posterURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	filename = strings.Replace(title, " ", "_", -1)
	filename += ".jpg"
	err = ioutil.WriteFile(filename, body, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}
