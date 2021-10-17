package helper

import "testing"

/*
before we created unit test, we must know the rules.
we have to include the `test` words in the name file
example hello_world_test.go

besides name file, the name of function still have rules,
the function in the first character must Capital. like `TestHelloWorld`,
and must have parameter (t *testing.T).

we can use `go test` to run all unit test.
if we want to see the parameter of function or what is function we run use `go test -v`
if w want to run the single function, we can use `go test -v -run TestHelloWorld` or `go test -v -run=TestHelloWorld` use equal (=)

we can run `go test -v ./...`, to running all unit testing
*/

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Aji")
	if result != "Hello Aji" {
		panic("Result is not Hello Aji")
	}
}

func TestHelloAji(t *testing.T)  {
	result := HelloWorld("Aji")
	if result != "Hello Aji" {
		panic("Result is not Hello Aji")
	}
}