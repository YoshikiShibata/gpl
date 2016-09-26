package params

import (
	"ch12/ex11/params"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

// Copyright © 2016 Yoshiki Shibata. All rights reserved.

//+ Exercise 12.12
// Added test for verify the original Unpack function.
func TestUnpack(t *testing.T) {
	type Data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	for _, test := range []struct {
		url  string
		data Data
	}{
		{`http://localhost:12345/search`, Data{nil, 10, false}},
		{`'http://localhost:12345/search?l=golang&l=programming`,
			Data{[]string{"golang", "programming"}, 10, false}},
		{`http://localhost:12345/search?l=golang&l=programming&max=100`,
			Data{[]string{"golang", "programming"}, 100, false}},
		{`http://localhost:12345/search?x=true&l=golang&l=programming`,
			Data{[]string{"golang", "programming"}, 10, true}},
	} {
		var data Data
		data.MaxResults = 10 // set default

		req, err := newRequest(test.url)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := params.Unpack(req, &data); err != nil {
			t.Errorf("Unpack: %v\n", err)
			continue
		}

		if !reflect.DeepEqual(data, test.data) {
			t.Errorf("%q => \n%+v but want %+v\n", test.url, data, test.data)
		}
	}
}

func TestUnpackExtension(t *testing.T) {
	type Data struct {
		// email allows either an empty string or a valid email address
		// such as XXXX@YYYY
		Mail string `http:"m,email"`

		// credit allow either an empty string or a valid 10 digits
		CreditNumber string `http:"c,credit"`

		// zip allows either an empty string or a valid 7 digits Japanese
		// zip code
		ZipCode string `http:"z,zip"`
	}

	for _, test := range []struct {
		url  string
		err  bool
		data Data
	}{
		{`http://localhost:12345/search`, false, Data{"", "", ""}},
		{`'http://localhost:12345/search?m=yoshiki.shibata@gmail.com`, false,
			Data{Mail: "yoshiki.shibata@gmail.com"}},
		{`'http://localhost:12345/search?m=@`, true, Data{}},
		{`'http://localhost:12345/search?m=yoshiki.shibata@`, true, Data{}},
		{`'http://localhost:12345/search?m=@gmail.com`, true, Data{}},
		{`'http://localhost:12345/search?c=1234567890`, false,
			Data{CreditNumber: "1234567890"}},
		{`'http://localhost:12345/search?c=1`, true, Data{}},
		{`'http://localhost:12345/search?c=12345678901`, true, Data{}},
		{`'http://localhost:12345/search?c=xyz`, true, Data{}},
		{`'http://localhost:12345/search?z=2270038`, false,
			Data{ZipCode: "2270038"}},
		{`'http://localhost:12345/search?z=2`, true, Data{}},
		{`'http://localhost:12345/search?z=22700380`, true, Data{}},
		{`'http://localhost:12345/search?z=xyz`, true, Data{}},
	} {
		var data Data

		req, err := newRequest(test.url)
		if err != nil {
			t.Errorf("Parse: %v\n", err)
			continue
		}

		if err := params.Unpack(req, &data); err != nil {
			if test.err == true {
				continue
			}

			t.Errorf("Unpack: %v\n", err)
			continue
		}

		if !reflect.DeepEqual(data, test.data) {
			t.Errorf("%q => \n%+v but want %+v\n", test.url, data, test.data)
		}
	}
}

func newRequest(rawurl string) (*http.Request, error) {
	var req http.Request
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	req.URL = url
	return &req, nil
}

//- Exercise 12.12
