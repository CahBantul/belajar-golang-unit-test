package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Nozami")
	if result != "hello Nozami" {
		t.Error("Expected 'hello Nozami' but got ", result)
		fmt.Println("test Hello world done")
	}
}

func TestHelloWorldAjitama(t *testing.T)  {
	result := HelloWorld("Ajitama")
	if result != "hello Ajitama" {
		t.Fatal("Expected 'hello Ajitama' but got ", result)
		fmt.Println("test Hello world done")
	}
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Ajitama")
	require.Equal(t, "hello Ajitama", result, "Expected 'hello Ajitama' but got %s", result)
	fmt.Println("tidak dieksekusi")
}

func TestHelloWorldAssertion(t *testing.T)  {
	result := HelloWorld("Ajitama")
	assert.Equal(t,  "hello Ajitama", result, "The two values are not equal.")
	fmt.Println("test Hello world with assertion done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Skipping test on linux because it is broken")
	}
}

func TestMain(m *testing.M)  {
	fmt.Println("Sebelum unit Test")
	m.Run()
	fmt.Println("Sesudah Unit Test")
}

func TestSubTest(t *testing.T)  {
	t.Run("Ajitama", func(t *testing.T) {
		res := HelloWorld("Ajitama")
		assert.Equal(t, res, "hello Ajitama")
	})

	t.Run("Kucing", func(t *testing.T) {
		res := HelloWorld("Kucing")
		assert.Equal(t, res, "hello Kucing")
	})
}

func TestHelloWorldTable(t *testing.T)  {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Nozami",
			request: "Nozami",
			expected: "hello Nozami",
		},
		{
			name: "Ajitama",
			request: "Ajitama",
			expected: "hello Ajitama",
		},
		{
			name: "FArdan",
			request: "FArdan",
			expected: "hello FArdan",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			got := HelloWorld(test.request)
			assert.Equal(t, got, test.expected)
		})
	}
}