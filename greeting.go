package greeting

import "fmt"

func Hello(name string) string {

	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
