package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

type Set struct {
	Vocabulary map[string]string
}

func (c *Set) Execute(input []string) (string, error) {
	src, value := input[0], input[1]
	c.Vocabulary[src] = value

	var respType resp.Type = &types.SimpleString{Content: "OK"}

	return respType.Encode(), nil
}
