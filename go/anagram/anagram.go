package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	out := make([]string, 0)
outer:
	for _, v := range candidates {
		in := make(map[int32]int)
		// create map to store each character as key and int value to be incremented
		for _, val := range strings.ToLower(subject) {
			if _, ok := in[val]; !ok {
				in[val] = 1
				continue
			}
			in[val] += 1
		}
		// validate if subject and candidate are equal length
		if len(subject) != len(v) {
			continue
		}
		// validate if subject and candidate are same
		if strings.ToLower(subject) == strings.ToLower(v) {
			continue
		}

		for _, sv := range strings.ToLower(v) {
			// validate if character is already present in the map
			if value, ok := in[sv]; ok {
				// validate if the weight of character is less than zero
				if value != 0 {
					in[sv] -= 1
					// if weight of the character is zero
					if in[sv] == 0 {
						// remove key and value
						delete(in, sv)
					}
				} else {
					continue
				}
			} else {
				continue outer
			}
		}
		// validate if the map is not empty
		if len(in) == 0 {
			out = append(out, v)
		}
	}
	return out
}
