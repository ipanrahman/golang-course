package unittest

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ipan")
	if result != "Hello, Ipan" {
		t.Error("Result must be Hello, Ipan")
	}
	fmt.Println("Testing TestHelloWorld done.")
}

func TestHelloAhmad(t *testing.T) {
	result := HelloWorld("Bukan Ahmad")
	if result != "Hello, Ahmad" {
		t.Fatal("Result must be 'Hello, Bukan Ahmad'")
	}
	fmt.Println("Testing TestHelloAhmad done.")
}
