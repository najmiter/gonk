package main

import "fmt"

var keywords = map[string][]string{
	"keywords": {
		"package", "import", "func",
	},
}

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
	lines := [][]string{}

	word := ""
	line := []string{}
	for _, char := range code {
		if space := IsSpace(char); space != Nil {
			switch space {
			case NewLine:
				if word == "" {
					continue
				}
				line = append(line, word)
				lines = append(lines, line)
				line = []string{}
				word = ""
			case Space:
				line = append(line, word)
				word = ""
			case Tab:
				continue
			}
		} else {
			word += string(char)
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}
