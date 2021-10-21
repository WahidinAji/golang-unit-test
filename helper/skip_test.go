package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSkip(t *testing.T)  {
	//if runtime.GOOS == "windows"{
	//	t.Skip("Can not run on Windows OS")
	//}

	result := HelloWorld("Aji")
	require.Equal(t,"Hello Aji",result,"Result must be 'Hello Aji' ")
	fmt.Println("TestHelloWorld With Skip Test Done") //if failed, this line is never called
}
