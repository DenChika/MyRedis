package types

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
)

type BulkString struct {
	Content string
}

func (t *BulkString) Encode() string {
	str := t.Content
	length := len(str)
	return fmt.Sprintf("$%d%s%s%s", length, resp.ClrfDelimeter, str, resp.ClrfDelimeter)
}
