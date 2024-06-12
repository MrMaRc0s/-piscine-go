package main

import (
    "github.com/01-edu/z01"
)

type point struct {
    x rune
    y rune
}

func setPoint(ptr *point) {
    ptr.x = 42
    ptr.y = 21
}

func printInt(n int) {
    if n == 0 {
        z01.PrintRune('0')
        return
    }
    
    var digits []rune
    isNegative := false
    
    if n < 0 {
        isNegative = true
        n = -n
    }
    
    for n > 0 {
        digits = append(digits, rune('0' + n % 10))
        n /= 10
    }
    
    if isNegative {
        z01.PrintRune('-')
    }
    
    for i := len(digits) - 1; i >= 0; i-- {
        z01.PrintRune(digits[i])
    }
}

func main() {
    points := &point{}

    setPoint(points)

    z01.PrintRune('x')
    z01.PrintRune(' ')
    z01.PrintRune('=')
    z01.PrintRune(' ')
    printInt(int(points.x))
    z01.PrintRune(',')
    z01.PrintRune(' ')
    z01.PrintRune('y')
    z01.PrintRune(' ')
    z01.PrintRune('=')
    z01.PrintRune(' ')
    printInt(int(points.y))
    z01.PrintRune('\n')
}