package types

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"strconv"
)

type Integer struct {
	Content string
}

func (t *Integer) Encode() string {
	v, _ := strconv.Atoi(t.Content)
	sign := '-'
	if v > 0 {
		sign = '+'
	}
	return fmt.Sprintf(":%b%s", sign, resp.ClrfDelimeter)
}
