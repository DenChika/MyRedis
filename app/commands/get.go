package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

type Get struct {
	Vocabulary *map[string]string
}

func (c *Get) Execute(input []string) (string, error) {
	src := input[0]
	value, ok := (*c.Vocabulary)[src]
	if !ok {
		return resp.EncodeEmpty(), nil
	}

	var respType resp.Type = &types.BulkString{Content: value}

	return respType.Encode(), nil
}
