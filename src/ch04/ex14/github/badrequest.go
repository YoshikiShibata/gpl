package github

import "fmt"

type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

type BadRequest struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

func (br *BadRequest) Error() string {
	return fmt.Sprintf("%s: %#v", br.Message, br.Errors)
}
