package types

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
)

type SimpleString struct {
	Content string
}

func (t *SimpleString) Encode() string {
	str := t.Content
	return fmt.Sprintf("+%s%s", str, resp.ClrfDelimeter)
}
