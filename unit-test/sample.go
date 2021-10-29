package unittest

import (
	"fmt"
)

func HelloWorld(v string) string {
	fmt.Println("Hello, ", v)
	return fmt.Sprintf("Hello, %s", v)
}
