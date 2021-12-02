package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	out := 0
	// If inputs has different length
	if len(a) != len(b){
		return out,fmt.Errorf("inputs length is not equal")
	}

	// Loop though either a or b because both are equal length
	// only interested on the index
	for i,_ := range a {
		// Validate if both characters are same
		if a[i] != b[i]{
			out++ // Increment output if characters are different
		}
	}
	return out, nil
}
