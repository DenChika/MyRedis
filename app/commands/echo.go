package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

type Echo struct{}

func (*Echo) Execute(input []string) (string, error) {

	if len(input) == 1 {
		bulk := &types.BulkString{Content: input[0]}
		return bulk.Encode(), nil
	}

	arr := make([]*types.BulkString, 0, len(input))
	for _, v := range input {
		arr = append(arr, &types.BulkString{Content: v})
	}

	return resp.EncodeArray(arr), nil
}
