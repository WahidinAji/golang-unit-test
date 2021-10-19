package helper

import (
	"fmt"
	"testing"
)

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


/*
if we will to fail result of testing, we can use Fail() and FailNow(),
Fail() will fail the unit test, but still run to execution the unit test. but in the last if already done, then the unit test will fail.
FailNow() will fail the unit test right now, without move on to execution the unit test.

Don't use panic for failed unit test.
 */

func TestHelloWorldFail(t *testing.T)  {
	result := HelloWorld("Aji")
	if result != "Hello Aji" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldAji Done")
}

func TestHelloWorldFailNow(t *testing.T)  {
	result := HelloWorld("Wahidin")
	if result != "Hello Wahidin" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldWahidin Done")
}

/*
and also, go had Fatal() and Error() to custom the information of your errors test
Fatal() called FailNow()

and so recommended to use Fatal() and Error(), because we can add information why our information is failed.
 */
// Error
func TestHelloWorldError(t *testing.T)  {
	result := HelloWorld("Aji")

	if result != "Hello Ajii"{
		t.Error("Result must be 'Hello Aji' ")
	}
	fmt.Println("TestHelloWorldError Done")
}// Fatal
func TestHelloWorldFatal(t *testing.T)  {
	result := HelloWorld("Wahidin")
	if result != "Hello Wahidinn"{
		t.Fatal("Result must be 'Hello Wahidin' ")
	}
	fmt.Println("TestHelloWorldFatal Done")
}