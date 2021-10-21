package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

//before after unit test,
func TestMain(m *testing.M)  {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

/*
to run Sub-Test u can use `go test -v -run=TestSubTest/Aji` or `-run=/Aji`. meanwhile, it's also running all subtest
in other test if it's had the same name (Aji) / (=TestSubTest/Aji).
 */
func TestSubTest(t *testing.T)  {
	//in the below of this line, this is a subTest.
	t.Run("Aji", func(t *testing.T) {
		result := HelloWorld("Aji")
		require.Equal(t,"Hello Aji",result,"Result must be 'Hello Aji' ")
	})
	t.Run("Success", func(t *testing.T) {
		result := HelloWorld("Success")
		require.Equal(t,"Hello Success",result,"Result must be 'Hello Success' ")
	})
	t.Run("Sub Fail", func(t *testing.T) {
		result := HelloWorld("Sub Fail")
		require.Equal(t,"Hello Sub Fail",result,"Result must be 'Hello Sub Fail' ")
	})
}
