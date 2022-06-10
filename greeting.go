package greeting

import "fmt"

func Hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}

func Bye(name string) string {
	msg := fmt.Sprintf("Oh, %v. Bye!", name)
	return msg
}
