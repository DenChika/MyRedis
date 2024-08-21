package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp/types"
)

type Ping struct{}

func (*Ping) Execute(_ []string) (string, error) {
	var respType resp.Type = &types.SimpleString{Content: "PONG"}
	return respType.Encode(), nil
}
