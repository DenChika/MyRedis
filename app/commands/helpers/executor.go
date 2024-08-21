package helpers

import (
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"strings"
)

type CliCommand interface {
	Execute([]string) (string, error)
}

func ExecuteCommand(str string) (string, error) {
	decoded := resp.Decode(str)

	name := strings.ToLower(decoded[0])
	var cmd CliCommand

	switch name {
	case "ping":
		cmd = &commands.Ping{}
		break
	case "echo":
		cmd = &commands.Echo{}
		break
	}

	return (cmd).Execute(decoded[1:])
}
