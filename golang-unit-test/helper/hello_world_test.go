package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAkbar(t *testing.T) {
	result := HelloWorld("Akbar")
	if result != "Hello Akbar" {
		t.Error("result must be 'Hello Akbar'")
	}
	fmt.Println("Test Akbar done")
}

func TestHelloWorldRizky(t *testing.T) {
	result := HelloWorld("Rizky")
	if result != "Hello Rizky" {
		t.Fatal("result must be 'Hello Rizky'")
	}

	fmt.Println("Test Rizky done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Akbar")
	assert.Equal(t, "Hello Akbar", result, "result must be 'Hello Akbar'")
	fmt.Println("Test w/ Assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Akbar")
	require.Equal(t, "Hello Akbar", result, "result must be 'Hello Akbar'")
	fmt.Println("Test w/ Require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" || runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		t.Skip("cannot run on unix like OS")
	}

	result := HelloWorld("Rizky")
	require.Equal(t, "Hello Rizky", result, "result must be 'Hello Rizky'")
}

func TestMain(m *testing.M) {
	fmt.Println("before unit test")
	m.Run()
	fmt.Println("after unit test")
}

func TestSubTest(t *testing.T) {
	t.Run("Akbar", func(t *testing.T) {
		result := HelloWorld("Akbar")
		require.Equal(t, "Hello Akbar", result, "result must be 'Hello Akbar'")
	})

	t.Run("Rizky", func(t *testing.T) {
		result := HelloWorld("Rizky")
		require.Equal(t, "Hello Rizky", result, "result must be 'Hello Rizky'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Muhammad",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "Rizky",
			request:  "Rizky",
			expected: "Hello Rizky",
		},
		{
			name:     "Akbar",
			request:  "Akbar",
			expected: "Hello Akbar",
		},
	}

	for _, test := range(tests){
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorldRizky(b *testing.B){
	for i := 0; i < b.N; i++ {
		fmt.Println("Rizky")
	}
}

func BenchmarkHelloWorldAkbar(b *testing.B){
	for i := 0; i < b.N; i++ {
		fmt.Println("Akbar")
	}
}

func BenchmarkSub(b *testing.B){
	b.Run("Rizky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Println("Rizky")
		}
	})
	b.Run("Akbar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fmt.Println("Akbar")
		}
	})
}

func BenchmarkTable(b *testing.B){
	tests := []struct{
		name string
		request string
	}{
		{
			name: "Muhammad",
			request: "Muhammad",
		},
		{
			name: "Rizky",
			request: "Rizky",
		},
		{
			name: "Akbar",
			request: "Akbar",
		},
		{
			name: "Muhammad Rizky Akbar",
			request: "Muhammad Rizky Akbar",
		},
	}

	for _, benchmark := range(tests){
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
