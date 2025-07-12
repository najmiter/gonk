package main

import (
	"errors"
	"strconv"
)

func Parse(code string) (tokenses [][]Token, _ error) {
	if code == "" {
		return nil, errors.New("Code is empty")
	}

	word := ""
	tokens := []Token{}

	for _, char := range code {
		if space := IsSpace(char); space != Nil {
			// get the kind
			kind, ok := kinds[word]
			if !ok {
				kind = KindPlain
			}

			// check for numeric values
			_, isNum := strconv.Atoi(word)
			if isNum == nil {
				kind = KindNumber
			}
			token := Token{name: word, kind: kind}
			switch space {
			case NewLine:
				if word == "" {
					continue
				}
				tokens = append(tokens, token)
				tokenses = append(tokenses, tokens)
				tokens = []Token{}
				word = ""
			case Space:
				tokens = append(tokens, token)
				word = ""
			case Tab:
				continue
			}
		} else {
			word += string(char)
		}
	}
	return tokenses, nil
}

func IsSpace(s rune) int {
	switch s {
	case 9: // "\t"
		return Tab
	case 10: // "\n"
		return NewLine
	case 13: // "\r"
		return WindowsThing
	case 32: // " "
		return Space
	}
	return Nil
}
