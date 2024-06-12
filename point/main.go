package main

import "unicode/utf8"

type point struct {
    x rune
    y rune
}

func setPoint(ptr *point) {
    ptr.x = 42
    ptr.y = 21
}

func main() {
    points := &point{}

    setPoint(points)

    // Convert runes to their string representations for printing
    xStr := string(points.x)
    yStr := string(points.y)

    // Calculate the width of the runes in bytes for demonstration
    xWidth := utf8.RuneLen(points.x)
    yWidth := utf8.RuneLen(points.y)

    // Print the results
    println("x =", xStr, "width in bytes:", xWidth)
    println("y =", yStr, "width in bytes:", yWidth)
}
