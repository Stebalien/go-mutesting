package example

import (
	"net/http"
)

func fooB() (a aType, b http.Header) {
	a, b = aType{}, http.Header{}

	return a, b
}
