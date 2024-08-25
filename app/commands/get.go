package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/resp/types"
	"sync"
)

type Get struct {
	Vocabulary map[string]string
	Mu         *sync.RWMutex
}

func (c *Get) Execute(input []string) (string, error) {
	src := input[0]
	c.Mu.RLock()
	value, ok := c.Vocabulary[src]
	c.Mu.RUnlock()
	if !ok {
		return resp.EncodeEmpty(), nil
	}

	var respType resp.Type = &types.BulkString{Content: value}

	return respType.Encode(), nil
}
