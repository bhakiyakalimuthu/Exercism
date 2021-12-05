package anagram

func Detect(subject string, candidates []string) []string {
	out := make([]string, 0)
	dict := make(map[int32]int)
	for _, v := range subject {
		if _, ok := dict[v]; !ok {
			dict[v] = 1
			continue
		}
		dict[v] += 1
	}

outer:
	for _, v := range candidates {
		in := make(map[int32]int)
		for dk, dv := range dict {
			in[dk] = dv
		}
		if len(subject) != len(v) {
			continue
		}
		for _, sv := range v {
			if value, ok := in[sv]; ok {
				if value != 0 {
					in[sv] -= 1
				} else {
					continue
				}
			} else {
				continue outer
			}
		}
		out = append(out, v)
	}
	return out
}
