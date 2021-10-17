package helper


/*
before we created unit test, we must know the rules,
we have to include the `test` words in the name file,
example hello_world_test.go.

besides name file, the name of function still have rules,
the function in the first character must Capital. like `TestHelloWorld`
and must have parameter (t *testing.T)
 */

func HelloWorld(name string) string  {
	return "Hello " + name
}
