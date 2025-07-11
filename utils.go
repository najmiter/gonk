package main

const (
	Tab = iota
	NewLine
	WindowsThing
	Space
	Nil
)

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
