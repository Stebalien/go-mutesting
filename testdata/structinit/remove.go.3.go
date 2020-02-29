// +build test

package example

type Some struct {
	a string
	B int
	c string
}

func removeFieldsFromStructInit() Some {
	aVal := "a value"
	cVal := "c value"

	some := Some{
		a: aVal,
		B: 5,
		c: cVal,
	}

	other := Some{}
	_ = aVal

	return some
}
