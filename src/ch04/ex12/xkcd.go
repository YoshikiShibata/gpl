// Copyright © 2016 Yoshiki Shibata

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var comics []*Comic

	num := 0
	notFoundCount := 0
	for {
		num++

		if exists(comicFilePath(num)) {
			fmt.Printf("Reading %5d ... ", num)
			comic := readComic(num)
			fmt.Printf("done\n")
			comics = append(comics, comic)
			continue
		}

		fmt.Printf("Fetching %5d ... ", num)
		comic, err, notFound := getComic(num)
		if err != nil {
			if notFound {
				fmt.Printf("Not Found - skipped\n")
				notFoundCount++
				if notFoundCount <= 10 {
					continue
				}

				fmt.Println("Probably no more comic")
				break

			}
			fmt.Printf("%v\n", err)
		}
		comics = append(comics, comic)
		fmt.Printf("done\n")
	}
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(fmt.Errorf("os.Stat: %v", err))
	}
	return true
}

const (
	xkcdURL      = "https://xkcd.com/"
	infoJson     = "info.0.json"
	xkcdDBDir    = "xkcd.db"
	xkcdInfoFile = xkcdDBDir + "/xkcd.%05d.info.json"
)

type Comic struct {
	Num        int
	Year       string
	Month      string
	Day        string
	Title      string
	Link       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
}

func getComic(num int) (*Comic, error, bool) {
	comicURL := fmt.Sprintf("%s/%d/%s", xkcdURL, num, infoJson)

	resp, err := http.Get(comicURL)
	if err != nil {
		return nil, err, false
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get Failed: %s", resp.Status),
			resp.StatusCode == http.StatusNotFound
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err, false
	}

	createDBDirectoryIfNecessary()
	err = ioutil.WriteFile(comicFilePath(num), jsonBytes, 0666)
	if err != nil {
		return nil, err, false
	}

	var comic Comic
	err = json.Unmarshal(jsonBytes, &comic)
	if err != nil {
		return nil, err, false
	}

	return &comic, nil, false
}

func readComic(num int) *Comic {
	bytes, err := ioutil.ReadFile(comicFilePath(num))
	if err != nil {
		panic(nil)
	}
	var comic Comic
	err = json.Unmarshal(bytes, &comic)
	if err != nil {
		panic(nil)
	}

	return &comic
}

func comicFilePath(num int) string {
	return fmt.Sprintf(xkcdInfoFile, num)
}

func createDBDirectoryIfNecessary() {
	fileInfo, err := os.Stat(xkcdDBDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(xkcdDBDir, 0777)
			if err != nil {
				panic(err)
			}
			return
		}
	}

	if !fileInfo.IsDir() {
		panic(fmt.Errorf("%s is not directory", xkcdDBDir))
	}
}
