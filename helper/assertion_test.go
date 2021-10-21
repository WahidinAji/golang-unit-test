package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

/*
assert called Fail()
require called FailNow()
 */

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Aji")
	assert.Equal(t, "Hello Aji",result,"Result must be 'Hello Aji' ")
	fmt.Println("TestHelloWorld With Assert Done")
}

//func TestHelloWorldAssertFail(t *testing.T)  {
//	result := HelloWorld("Aji Aja")
//	assert.Equal(t, "Hello Aji",result,"Result must be 'Hello Aji' ")
//	fmt.Println("TestHelloWorld With Assert Fail")
//}

//require
func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Aji")
	require.Equal(t,"Hello Aji",result,"Result must be 'Hello Aji' ")
	fmt.Println("TestHelloWorld With Require Done") //if failed, this line is never called
}
//
//func TestHelloWorldRequireFail(t *testing.T)  {
//	result := HelloWorld("Aji Aja")
//	require.Equal(t, "Hello Aji",result,"Result must be 'Hello Aji' ")
//	fmt.Println("TestHelloWorld With Require Fail") //if failed, this line is never called
//}
