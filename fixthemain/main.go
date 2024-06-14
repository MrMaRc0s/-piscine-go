package piscine

import "github.com/01-edu/z01"

// Define constants for door states
const (
	OPEN  = true
	CLOSE = false
)

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	var ptr []rune = []rune("Door Opening...")
	for i := 0; i < len(ptr); i++ {
		z01.PrintRune(ptr[i])
	}
	z01.PrintRune('\n')
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	var ptr []rune = []rune("Door Closing...")
	for i := 0; i < len(ptr); i++ {
		z01.PrintRune(ptr[i])
	}
	z01.PrintRune('\n')
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(door *Door) bool {
	var ptr []rune = []rune("Is the Door opened?")
	for i := 0; i < len(ptr); i++ {
		z01.PrintRune(ptr[i])
	}
	z01.PrintRune('\n')
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	var ptr []rune = []rune("Is the Door closed?")
	for i := 0; i < len(ptr); i++ {
		z01.PrintRune(ptr[i])
	}
	z01.PrintRune('\n')
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
