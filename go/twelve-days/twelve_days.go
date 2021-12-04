package twelve

import "bytes"

const numOfVerses = 12

func Verse(i int) string {

	m := map[string]string{
		"first":    "a Partridge in a Pear Tree.",
		"second":   "two Turtle Doves, ",
		"third":    "three French Hens, ",
		"fourth":   "four Calling Birds, ",
		"fifth":    "five Gold Rings, ",
		"sixth":    "six Geese-a-Laying, ",
		"seventh":  "seven Swans-a-Swimming, ",
		"eighth":   "eight Maids-a-Milking, ",
		"ninth":    "nine Ladies Dancing, ",
		"tenth":    "ten Lords-a-Leaping, ",
		"eleventh": "eleven Pipers Piping, ",
		"twelfth":  "twelve Drummers Drumming, ",
	}

	var n = map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
		13: "On the ",
		14: " day of Christmas my true love gave to me: ",
	}
	var in, out bytes.Buffer
	out.WriteString(n[13] + n[i] + n[14])
	for y := i; y > 0; y-- {
		if i != 1 && y == 1 {
			in.WriteString("and ")
		}
		in.WriteString(m[n[y]])
	}
	out.WriteString(in.String())
	return out.String()
}

func Song() string {
	var out bytes.Buffer
	for x := 1; x <= numOfVerses; x++ {
		out.WriteString(Verse(x))
		if x != numOfVerses {
			out.WriteString("\n")
		}
	}
	return out.String()
}
