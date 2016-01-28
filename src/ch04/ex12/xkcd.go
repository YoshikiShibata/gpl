// Copyright © 2016 Yoshiki Shibata

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

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

type cacheResult struct {
	num   int
	comic *Comic
	ok    bool
}

const concurrencyLevel = 10 // Concurrency Level and Max Not Found
const lastComicNumber = 1635

var vFlag = flag.Bool("v", false, "verbose message")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No word specified\n")
		os.Exit(1)
	}
	fmt.Printf("Building indexes ... ")
	comics := fetchAllComics()
	fmt.Printf("Done\n")
	findWords(comics, flag.Args())
}

func findWords(comics map[int]*Comic, words []string) {
topLoop:
	for i := 1; i <= lastComicNumber; i++ {
		comic := comics[i]
		if comic == nil {
			continue
		}

		for _, word := range words {
			if !strings.Contains(comic.Transcript, word) {
				continue topLoop
			}
		}
		fmt.Printf("\n==== Matched =====\nURL: %s\n", xkcdURL+strconv.Itoa(i))
		fmt.Printf("Transcript: %s\n", comic.Transcript)
	}
}

var wg sync.WaitGroup

func fetchAllComics() map[int]*Comic {
	var comics = make(map[int]*Comic)
	var results = make(chan *cacheResult, concurrencyLevel)

	num := 0
	num++
	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go cacheComic(results, num)
		num++
	}

	chanelClosed := false
	for result := range results {
		if result.ok {
			vPrintf("caching %5d ... done\n", result.num)
			comics[result.num] = result.comic
		}
		wg.Add(1)
		go cacheComic(results, num)
		num++
		if num > lastComicNumber && !chanelClosed {
			vPrintf("Waiting ... \n")
			wg.Wait()
			close(results)
			chanelClosed = true
			vPrintf("Done \n")
		}
	}
	return comics
}

func vPrintf(format string, args ...interface{}) {
	if *vFlag {
		fmt.Printf(format, args...)
	}
}

func cacheComic(comics chan<- *cacheResult, num int) {
	defer wg.Done()

	if exists(comicFilePath(num)) {
		vPrintf("Reading %5d ...  \n", num)
		comic := readComic(num)
		comics <- &cacheResult{num, comic, true}
		return
	}

	vPrintf("Fetching %5d ... \n", num)
	comic, err, notFound := getComic(num)
	if err != nil {
		if notFound {
			vPrintf("%5d ... not found\n", num)
			comics <- &cacheResult{num, nil, false}
			return
		}
		vPrintf("%v\n", err)
		comics <- &cacheResult{num, nil, false}
		return
	}
	comics <- &cacheResult{num, comic, true}
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
