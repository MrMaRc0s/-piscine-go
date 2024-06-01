package piscine

func BasicAtoi(s string) int {
    result := 0

    for i := 0; i < len(s); i++ {
        // Convert the character to its numeric value by subtracting '0'
        digit := int(s[i] - '0')
        result = result*10 + digit
    }

    return result
}
