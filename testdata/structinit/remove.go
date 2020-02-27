// +build test

package example

type Some struct {
	a string
	B int
}

func removeFieldsFromStructInit() Some {
	some := Some{
		a: "a value",
		B: 5,
	}

	return some
}
