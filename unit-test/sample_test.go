package unittest

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ipan")
	if result != "Hello, Ipan" {
		panic("Result is not ipan")
	}
}
