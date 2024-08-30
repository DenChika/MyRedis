package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/commands"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
	"strings"
	"sync"
)

type CliCommand interface {
	Execute([]string) (string, error)
}

var vocabulary map[string]string = make(map[string]string)

type Executor struct {
	vocabulary map[string]string
	mu         *sync.RWMutex
}

func GetOrCreateExecutor() *Executor {
	return &Executor{vocabulary: vocabulary, mu: &sync.RWMutex{}}
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
		cmd = &commands.Set{Vocabulary: e.vocabulary, Mu: e.mu}
		break
	case "get":
		cmd = &commands.Get{Vocabulary: e.vocabulary, Mu: e.mu}
		break
	}

	return (cmd).Execute(decoded[1:])
}
