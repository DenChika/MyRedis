package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"strings"
)

type CliCommand interface {
	Execute([]string) (string, error)
}

type Executor struct{}

func GetExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(str string) (string, error) {
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
	case "set":
		cmd = &commands.Set{Vocabulary: Storage.Vocabulary, Mu: Storage.Mu}
		break
	case "get":
		cmd = &commands.Get{Vocabulary: Storage.Vocabulary, Mu: Storage.Mu}
		break
	}

	return (cmd).Execute(decoded[1:])
}
