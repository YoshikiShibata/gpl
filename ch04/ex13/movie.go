// Copyright © 2016 Yoshiki Shibata

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type errorJSON struct {
	Response string
	Error    string
}

type movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string
	ImdbVotes  string
	ImdbID     string
	Type       string
	Response   string
}

func (m *movie) String() string {
	json, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(json)
}

const queryTemplate = "http://www.omdbapi.com/?t=%s&y=&plot=short&r=json"

func getMovie(title string) (*movie, error) {
	queryURL := fmt.Sprintf(queryTemplate, url.QueryEscape(title))

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var errJSON errorJSON
	err = json.Unmarshal(body, &errJSON)
	if err != nil {
		return nil, err
	}

	if errJSON.Response == "False" {
		return nil, fmt.Errorf("%s", errJSON.Error)
	}

	var mv movie
	err = json.Unmarshal(body, &mv)
	if err != nil {
		return nil, err
	}

	return &mv, nil
}
