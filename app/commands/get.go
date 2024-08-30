package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/lib/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/resp/types"
	"sync"
	"time"
)

type Get struct {
	Vocabulary map[string]commands.ExpiryValue
	Mu         *sync.RWMutex
}

func (c *Get) Execute(input []string) (string, error) {
	src := input[0]

	c.Mu.RLock()

	exp, ok := c.Vocabulary[src]
	if time.Now().After(exp.ValidUntil) || !ok {
		c.Mu.RUnlock()
		return resp.EncodeEmpty(), nil
	}

	c.Mu.RUnlock()

	value := exp.Value
	var respType resp.Type = &types.BulkString{Content: value}

	return respType.Encode(), nil
}
