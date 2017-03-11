package excercises

func ReverseCharsInPlace(value []byte) {
	// Hello World
	// dlorw olleh
	j := len(value) - 1
	for i := 0; i < len(value)/2; i++ {
		c := value[j]
		value[j] = value[i]
		value[i] = c
		j--
	}
}
