package main

import "fmt"

func main() {
	code := `
	package main
	import "fmt"
	func main() {
		var number := 123
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
