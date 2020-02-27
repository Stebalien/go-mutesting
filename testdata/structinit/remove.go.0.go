// +build test

package example

type Some struct {
	a string
	B int
}

func removeFieldsFromStructInit() Some {
	some := Some{
		B: 5,
	}

	return some
}
