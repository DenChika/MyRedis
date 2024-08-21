package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

type Get struct{}

func (*Get) Execute(input []string) (string, error) {
	src := input[0]
	value, ok := SetMemoryData[src]
	if !ok {
		return resp.EncodeEmpty(), nil
	}

	var respType resp.Type = &types.BulkString{Content: value}

	return respType.Encode(), nil
}
