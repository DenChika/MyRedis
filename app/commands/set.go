package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
	"sync"
)

type Set struct {
	Vocabulary map[string]string
	Mu         *sync.Mutex
}

func (c *Set) Execute(input []string) (string, error) {
	src, value := input[0], input[1]

	c.Mu.Lock()
	c.Vocabulary[src] = value
	c.Mu.Unlock()

	var respType resp.Type = &types.SimpleString{Content: "OK"}

	return respType.Encode(), nil
}
