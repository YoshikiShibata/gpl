package params

import "testing"

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func TestPack(t *testing.T) {
	for _, test := range []struct {
		d      data
		params string
	}{{data{Labels: []string{"golang"}, MaxResults: 10, Exact: true},
		"l=golang&max=10&x=true"},
	} {
		params, err := Pack(&test.d)
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}

		if params != test.params {
			t.Error("params = %q, but want %q\n", params, test.params)
		}
	}
}
