package unittest

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ipan")
	if result != "Hello, Ipan" {
		t.Fail()
	}
	fmt.Println("Testing TestHelloWorld done.")
}

func TestHelloAhmad(t *testing.T) {
	result := HelloWorld("Bukan Ahmad")
	if result != "Hello, Ahmad" {
		t.FailNow()
	}
	fmt.Println("Testing TestHelloAhmad done.")
}
