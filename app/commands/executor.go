package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"strings"
)

type CliCommand interface {
	Execute([]string) (string, error)
}

func Execute(str string) (string, error) {
	decoded := resp.Decode(str)

	name := strings.ToLower(decoded[0])
	var cmd CliCommand

	switch name {
	case "ping":
		cmd = &Ping{}
		break
	case "echo":
		cmd = &Echo{}
		break
	case "set":
		cmd = &Set{}
		break
	case "get":
		cmd = &Get{}
		break
	}

	return (cmd).Execute(decoded[1:])
}
