package twelve

import "bytes"

type letterList []string

func Verse(i int) string {

	//On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.
	//
	//	On the second day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree.

	//On the twelfth day of Christmas my true love gave to me:
	//twelve Drummers Drumming,
	//eleven Pipers Piping,
	//ten Lords-a-Leaping,
	//nine Ladies Dancing,
	//eight Maids-a-Milking,
	//seven Swans-a-Swimming,
	//six Geese-a-Laying,
	//five Gold Rings,
	//four Calling Birds,
	//three French Hens,
	//two Turtle Doves, and a Partridge in a Pear Tree.

	m := map[string]string{
		"first":    "a Partridge in a Pear Tree",
		"second":   "two Turtle Doves,",
		"third":    "three French Hens,",
		"fourth":   "four Calling Birds,",
		"fifth":    "five Gold Rings,",
		"sixth":    "six Geese-a-Laying,",
		"seventh":  "seven Swans-a-Swimming,",
		"eighth":   "eight Maids-a-Milking,",
		"ninth":    "nine Ladies Dancing,",
		"tenth":    "ten Lords-a-Leaping,",
		"eleventh": "eleven Pipers Piping,",
		"twelfth":  "twelve Drummers Drumming,",
	}

	var NumberToWord = map[int]string{
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
	var out bytes.Buffer
	out.WriteString(NumberToWord[13] + NumberToWord[i] + NumberToWord[14])
	for x := 1; x <= i; x++ {
		var in bytes.Buffer
		for y := x; y > 0; y-- {
			if x != 1 && y == 1 {
				in.WriteString("and ")
			}
			in.WriteString(m[NumberToWord[y]])
		}

		out.WriteString(in.String())
		if i == 1 {
			out.WriteString(".")
		}
		if x != i {
			out.WriteString("\n")
		}

	}
	return out.String()
}

func Song() string {
	var out bytes.Buffer
	for x := 1; x <= 12; x++ {
		out.WriteString(Verse(x))
	}
	return out.String()
}
