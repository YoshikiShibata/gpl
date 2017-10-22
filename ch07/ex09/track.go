// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracksData = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Close to You", "Carpenters", "Close to You", 1970, length("0m00s")},
	{"Top of the World", "Carpenters", "A Song for You", 1972, length("0m00s")},
	{"Yesterday Once More", "Carpenters", "Now and Then", 1973, length("0m00s")},
	{"案山子", "さだまさし", "天晴", 2013, length("0m00s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 列幅を計算して表を印字する
}

type TracksTable struct {
	Tracks []*Track
}

func printTracksHTML(out io.Writer, tracks []*Track, keys []string) {
	titleFunc := func() template.HTML { return createQueryLink(keys, "Title") }
	artistFunc := func() template.HTML { return createQueryLink(keys, "Artist") }
	albumFunc := func() template.HTML { return createQueryLink(keys, "Album") }
	yearFunc := func() template.HTML { return createQueryLink(keys, "Year") }
	lengthFunc := func() template.HTML { return createQueryLink(keys, "Length") }

	funcMap := template.FuncMap{
		"title":  titleFunc,
		"artist": artistFunc,
		"album":  albumFunc,
		"year":   yearFunc,
		"length": lengthFunc}

	err := template.Must(template.New("tracktable").
		Funcs(funcMap).
		Parse(`
		<html>
		<head>
		<meta http-equiv="Content-Type" conntent="text/html; charset=utf-8">
		<title>My Tracks</title>
		</head>
		</body>
		<table border="5" rules="all" cellpadding="5">
		<tr style='text-align: left'>
			<th>{{title}}</th>
			<th>{{artist}}</th>
			<th>{{album}}</th>
			<th>{{year}}</th>
			<th>{{length}}</th>
		</tr>
		{{range .Tracks}}
		<tr>
			<td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
			<td>{{.Length}}</td>
		</tr>
		{{end}}
		</table>
		</body>
		</html>
	`)).Execute(out, &TracksTable{tracks})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func createQueryLink(keys []string, name string) template.HTML {
	updatedKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		if strings.ToUpper(key) != strings.ToUpper(name) {
			updatedKeys = append(updatedKeys, key)
		}
	}
	updatedKeys = append(updatedKeys, name)
	queryLink := fmt.Sprintf("<a href=\"?sort=%s\">%s</a>",
		strings.Join(updatedKeys, ","), name)
	return template.HTML(queryLink)
}
