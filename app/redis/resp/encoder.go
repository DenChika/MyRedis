package resp

import (
	"bytes"
	"fmt"
)

type Type interface {
	Encode() string
}

const (
	ClrfDelimeter = "\r\n"
)

func EncodeArray[T Type](arr []T) string {
	length := len(arr)
	arrMark := fmt.Sprintf("*%d%s", length, ClrfDelimeter)

	buf := bytes.NewBuffer([]byte(arrMark))

	for _, v := range arr {
		encoded := fmt.Sprintf("%v%s", v.Encode(), ClrfDelimeter)
		buf.WriteString(encoded)
	}
	return buf.String()
}
