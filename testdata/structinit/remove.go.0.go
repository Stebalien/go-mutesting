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

		B: 5,
		c: cVal,
	}
	_ = aVal

	var other Some

	other = Some{
		a: aVal,
	}

	return some
}
