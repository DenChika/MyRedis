package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"github.com/codecrafters-io/redis-starter-go/app/resp/types"
)

type Ping struct{}

func (*Ping) Execute(_ []string) (string, error) {
	var respType resp.Type = &types.SimpleString{Content: "PONG"}
	return respType.Encode(), nil
}
