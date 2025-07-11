package main

import "fmt"

func main() {
	code := `
	package main
	import "fmt"
	func main() {
		fmt.Println("Hello, World!")
		fmt.Println("This is a Go program.")
		fmt.Println("It prints multiple lines of text.")
		fmt.Println("Goodbye!")
	}`
	tokenses, err := Parse(code)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, tokens := range tokenses {
		for _, token := range tokens {
			fmt.Println(token)
		}
		fmt.Println("--------------")
	}

}
