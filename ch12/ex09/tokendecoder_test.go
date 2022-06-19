// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package sexpr

import (
	"fmt"
	"reflect"
	"testing"
)

// + Exercise 12.9
func TestTokenDecoder(t *testing.T) {
	for _, test := range []struct {
		sexpr  string
		tokens []Token
	}{
		{"( )", []Token{StartList{}, EndList{}}},
		{"(ABC)", []Token{StartList{},
			Symbol{"ABC"},
			EndList{}}},
		{`(ABC "DEF")`, []Token{StartList{},
			Symbol{"ABC"},
			String{"DEF"},
			EndList{}}},
		{`(ABC nil)`, []Token{StartList{},
			Symbol{"ABC"},
			Symbol{"nil"},
			EndList{}}},
		{`("ABC" nil)`, []Token{StartList{},
			String{"ABC"},
			Symbol{"nil"},
			EndList{}}},
		{`("ABC" 10)`, []Token{StartList{},
			String{"ABC"},
			Int{10},
			EndList{}}},
		{`(ABC (x 10))`, []Token{StartList{},
			Symbol{"ABC"},
			StartList{},
			Symbol{"x"},
			Int{10},
			EndList{},
			EndList{}}},
	} {
		d := NewDecoder([]byte(test.sexpr))
		for _, token := range test.tokens {
			next := d.Token()
			if !reflect.DeepEqual(next, token) {
				t.Fatal(fmt.Errorf("%#v, but want %#v", next, token))
			}
		}
	}
}

//- Exercise 12.9
