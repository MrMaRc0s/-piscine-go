// Quest4
package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	root := 0
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			root = i
			break
		}
	}
	return root
}
