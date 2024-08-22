package commands

import (
	"github.com/codecrafters-io/redis-starter-go/app/redis/resp"
	"strings"
	"sync"
)

type CliCommand interface {
	Execute([]string) (string, error)
}

type Executor struct {
	Vocabulary map[string]string
	Mu         *sync.Mutex
}

func NewExecutor() *Executor {
	return &Executor{Vocabulary: make(map[string]string), Mu: &sync.Mutex{}}
}

func (e *Executor) Execute(str string) (string, error) {
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
		cmd = &Set{Vocabulary: e.Vocabulary, Mu: e.Mu}
		break
	case "get":
		cmd = &Get{e.Vocabulary}
		break
	}

	return (cmd).Execute(decoded[1:])
}
