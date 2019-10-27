// Copyright Some Company Corp.
// All Rights Reserved
package toto


// here are the imports
import (
	"fmt"
)

// Hello to say hello
func HelloBis(test string) string{
	fmt.Println("hello world")
	// we return string here
	return test

}

func ExampleHelloBis(test string) string{
	fmt.Println("hello world")
	return test
    // Output: test

}

// Hello to say hello
func Hello() {
	fmt.Println("hello world")

}

