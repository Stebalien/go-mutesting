package example

import (
	"net/http"
)

type aType struct{}

func fooA() (a aType, b http.Header) {
	_, _, _ = a, b, http.Header{}

	return a, b
}
