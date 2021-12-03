package raindrops

import (
	"bytes"
	"fmt"
	"strconv"
)

func Convert1(number int) string {
	var out string
	var noCaseMatch int
	if number % 3 == 0{
		out += "Pling"
		noCaseMatch++
	}
	if number % 5 == 0 {
		out += "Plang"
		noCaseMatch++
	}
	if number % 7 == 0 {
		out += "Plong"
		noCaseMatch++
	}
	if noCaseMatch ==0 {
		return fmt.Sprintf("%d",number)
	}

	return out
}

func Convert(number int) string {

	in := []struct{
		num int
		word string
	}{
		{3,"Pling"}, {5, "Plang"}, {7, "Plong"},
	}
	var out bytes.Buffer
	for _, item := range in {
		if number % item.num == 0 {
			out.WriteString(item.word)
		}
	}
	if out.String() == ""{
		out.WriteString(strconv.Itoa(number))
	}
	return out.String()
}
