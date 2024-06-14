package piscine

import (
	"fmt"
)

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
		fmt.Printf("%c", r)
	}
	fmt.Println()
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("Is the Door opened?")
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?")
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
