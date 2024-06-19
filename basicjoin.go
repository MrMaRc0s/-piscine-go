// Quest5
package piscine

func BasicJoin(elems []string) string {
	var str string
	for i := 0; i < len(elems); i++ {
		str += string(elems[i])
	}
	return str
}
