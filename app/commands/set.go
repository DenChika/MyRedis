package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

var SetMemoryData = make(map[string]string)

type Set struct{}

func (*Set) Execute(input []string) (string, error) {
	src, value := input[0], input[1]
	SetMemoryData[src] = value

	var respType resp.Type = &types.SimpleString{Content: "OK"}

	return respType.Encode(), nil
}
