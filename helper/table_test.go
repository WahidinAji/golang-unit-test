package helper

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Table test is where we provide slice data a have parameter content and result expectation of unit test.
and then we iterate the slice using sub-test.
 */

func TestHelloWorldTable(t *testing.T)  {
	tests := []struct{
		name		string 	//name of sub-test
		request		string 	//what's ur request?
		expected	string	//expectation
	}{
		{
			name: "Aji",
			request: "Aji",
			expected: "Hello Aji",
		},
		{
			name: "Wahidin",
			request: "Wahidin",
			expected: "Hello Wahidin",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
