package resp

import (
	"fmt"
	"strings"
)

func clrfSplit(str string) []string {
	return strings.Split(str, ClrfDelimeter)
}

func Decode(str string) []string {
	split := clrfSplit(str)
	fmt.Println(split)

	result := make([]string, 0, len(split))
	for i := 0; i < len(split); i++ {
		switch split[i][0] {
		case '$':
			result = append(result, split[i+1])
			break
		case '+':
			result = append(result, split[i][1:])
		default:
			break
		}
	}

	return result
}
